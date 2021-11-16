package rpc_server

import (
	"fmt"
	"log"
	"net"

	"github.com/evenlab/go-kit/context"
	"github.com/jkrus/keep_connection/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"keep_connection_server/config"
	"keep_connection_server/rpc_server/handlers"
)

func RPCServer(cfg *config.Config) *grpc.Server {
	var (
		// please add in alphabetical order
		pongHandler = handlers.Pong{ConnIdle: cfg.MaxConnectionIdle}
	)

	gRPCServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: cfg.MaxConnectionIdle,
		}),
	)

	// please add in alphabetical order
	pb.RegisterPingPongServer(gRPCServer, &pongHandler)

	return gRPCServer
}

// Start starts the gRPC server.
func Start(
	ctx context.Context,
	cfg *config.Config,
	server *grpc.Server) {
	// start the server

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to listen: %v", err))
	}

	createContextHandler(ctx.WithCancelWait(), server)

	log.Printf("Listen on: %s:%v and serve...", cfg.Host, cfg.Port)

	if err = server.Serve(listener); err != nil {
		log.Fatal(fmt.Errorf("failed to serve: %v", err))
	}
}

// createContextHandler creates a context handler goroutine.
func createContextHandler(ctx context.Context, server *grpc.Server) {
	go func() {
		<-ctx.Done()
		defer ctx.WgDone()

		log.Println("RPC server shutting down...")
		server.Stop()
		log.Println("RPC server shutdown complete.")
	}()
}
