package service

import (
	"log"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"
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

// Ping implements Ping interface.
func (s *pingService) Ping() error {
	// TODO implement protocol here.

	return nil
}

// Start implements Ping interface.
func (s *pingService) Start() error {
	log.Println("Starts admin service...")

	s.createContextHandler()

	return nil
}

// Stop implements Ping interface.
func (s *pingService) Stop() error {
	log.Println("Stops admin service...")
	s.ctx.Cancel()
	s.ctx.WgDone()
	log.Println("Ping service stop complete.")

	return nil
}

// createContextHandler creates a context handler goroutine.
func (s *pingService) createContextHandler() {
	go func() {
		<-s.ctx.Done()
		_ = s.Stop()
	}()
}
