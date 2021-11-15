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
	// pingService implements Ping interface.
	pingService struct {
		ctx context.Context
		dic *di.Container
	}
)

var (
	// Make sure pingService implements Ping interface.
	_ Ping = (*pingService)(nil)
)

// NewPingService returns Ping interface.
func NewPingService(ctx context.Context, dic *di.Container) Ping {
	return &pingService{
		// use context with cancel and wait
		// to gracefully stop pingService
		ctx: ctx.WithCancelWait(),
		dic: dic,
	}
}

// keepConnection ...
func (s *pingService) keepConnection() error {
	var cfg *config.Config
	if err := s.dic.Resolve(&cfg); err != nil {
		return nil
	}

	pingTime := time.Duration(cfg.MaxConnectionIdle/2) * 1000 * 1000 * 1000
	timeOut := time.Duration(cfg.TimeOut) * 1000 * 1000 * 1000
	timeOutMsg := time.Duration(cfg.MessageTimeOut) * 1000 * 1000

	pingTimer := time.NewTimer(pingTime)
	timeOutTimer := time.NewTimer(timeOut)
	msgTimeOut := time.NewTimer(timeOutMsg)
	timeOutTimer.Stop()
	msgTimeOut.Stop()
	defer func() {
		pingTimer.Stop()
		timeOutTimer.Stop()
		msgTimeOut.Stop()
	}()

	var count int
	log.Println("1")
	for {
		select {
		case <-pingTimer.C:
			log.Println("2")
			ack, err := rpc.RenewalRequest()
			if err != nil {
				log.Println(err.Error())
			}
			if !ack {
				msgTimeOut.Reset(timeOutMsg)
			}
			pingTimer.Reset(pingTime)

		case <-msgTimeOut.C:
			log.Println("3")
			ack, err := rpc.RenewalRequest()
			if err != nil {
				log.Println(err.Error())
			}
			if !ack {
				count++
				if count < 2 {
					timeOutTimer.Reset(timeOutMsg)
					continue
				}
				count = 0
				msgTimeOut.Reset(timeOut)
			}

		case <-timeOutTimer.C:
			log.Println("4")
			ack, err := rpc.RenewalRequest()
			if err != nil {
				log.Println(err.Error())
			}
			if !ack {
				msgTimeOut.Reset(timeOutMsg)
			}
		}
	}

	return nil
}

// Start implements Ping interface.
func (s *pingService) Start() error {
	log.Println("Starts keep connection service...")

	s.createContextHandler()

	go s.keepConnection()

	return nil
}

// Stop implements Ping interface.
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
