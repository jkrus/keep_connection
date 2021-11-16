package service

import (
	"log"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"
)

type (
	// pongService implements Pong interface.
	pongService struct {
		ctx context.Context
		dic *di.Container
	}
)

var (
	// Make sure pongService implements Pong interface.
	_ Pong = (*pongService)(nil)
)

// NewKeepConnectService returns Pong interface.
func NewKeepConnectService(ctx context.Context, dic *di.Container) Pong {
	return &pongService{
		// use context with cancel and wait
		// to gracefully stop pongService
		ctx: ctx.WithCancelWait(),
		dic: dic,
	}
}

// Start implements Pong interface.
func (s *pongService) Start() error {
	log.Println("Starts pong service...")

	s.createContextHandler()

	return nil
}

// Stop implements Pong interface.
func (s *pongService) Stop() error {
	log.Println("Stops pong service...")
	s.ctx.Cancel()
	s.ctx.WgDone()
	log.Println("keepConnection service stop complete.")

	return nil
}

// createContextHandler creates a context handler goroutine.
func (s *pongService) createContextHandler() {
	go func() {
		<-s.ctx.Done()
		_ = s.Stop()
	}()
}
