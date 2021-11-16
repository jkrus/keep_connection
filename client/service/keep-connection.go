package service

import (
	"log"
	"time"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"

	"keep_connection_client/config"
	"keep_connection_client/rpc"
)

type (
	// pingService implements KeepConnect interface.
	pingService struct {
		ctx context.Context
		dic *di.Container
	}
)

var (
	// Make sure pingService implements KeepConnect interface.
	_ KeepConnect = (*pingService)(nil)
)

// NewKeepConnectService returns KeepConnect interface.
func NewKeepConnectService(ctx context.Context, dic *di.Container) KeepConnect {
	return &pingService{
		// use context with cancel and wait
		// to gracefully stop pingService
		ctx: ctx.WithCancelWait(),
		dic: dic,
	}
}

// keepConnection ...
func (s *pingService) keepConnection() {
	var cfg *config.Config
	if err := s.dic.Resolve(&cfg); err != nil {
		log.Println(err)
		return
	}

	pingTimer := time.NewTimer(cfg.MaxConnectionIdle / 2)
	timeOutTimer := time.NewTimer(cfg.TimeOut)
	msgTimeOut := time.NewTimer(cfg.MessageTimeOut)

	timeOutTimer.Stop()
	msgTimeOut.Stop()
	defer func() {
		pingTimer.Stop()
		timeOutTimer.Stop()
		msgTimeOut.Stop()
	}()

	var count int
	var attempts int
	for {
		if attempts > cfg.LimitAttempts {
			log.Println("attempts is end")
			break
		}
		select {
		case <-pingTimer.C:
			log.Println("ping")
			ack, err := rpc.RenewalRequest()
			if err != nil {
				log.Println(err.Error())
			}
			if !ack {
				msgTimeOut.Reset(cfg.MessageTimeOut)
				continue
			}
			pingTimer.Reset(cfg.MaxConnectionIdle / 2)

		case <-msgTimeOut.C:
			log.Println("ping timeOut")
			ack, err := rpc.RenewalRequest()
			if err != nil {
				log.Println(err.Error())
			}
			if !ack {
				count++
				if count < 2 {
					timeOutTimer.Reset(cfg.MessageTimeOut)
					continue
				}
				count = 0
				msgTimeOut.Reset(cfg.TimeOut)
				attempts++
				log.Println("attempts = ", attempts)
				continue
			}
			attempts = 0
			pingTimer.Reset(cfg.MaxConnectionIdle / 2)
		case <-timeOutTimer.C:
			log.Println("ping serial")
			ack, err := rpc.RenewalRequest()
			if err != nil {
				log.Println(err.Error())
			}
			if !ack {
				msgTimeOut.Reset(cfg.MessageTimeOut)
			}
		}
	}

	return
}

// Start implements KeepConnect interface.
func (s *pingService) Start() error {
	log.Println("Starts keep connection service...")

	s.createContextHandler()

	go s.keepConnection()

	return nil
}

// Stop implements KeepConnect interface.
func (s *pingService) Stop() error {
	log.Println("Stops keep connection service...")
	s.ctx.Cancel()
	s.ctx.WgDone()
	log.Println("Keep Connection service stop complete.")

	return nil
}

// createContextHandler creates a context handler goroutine.
func (s *pingService) createContextHandler() {
	go func() {
		<-s.ctx.Done()
		_ = s.Stop()
	}()
}
