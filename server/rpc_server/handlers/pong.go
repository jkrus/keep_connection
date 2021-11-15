package handlers

import (
	"context"
	"log"

	"github.com/jkrus/keep_connection/pb"
)

type Pong struct {
	pb.UnsafePingPongServer
}

var _ pb.PingPongServer = (*Pong)(nil)

// PingMessage implements method PingPongServer.PingMessage.
func (p *Pong) PingMessage(_ context.Context, in *pb.PingPongRequest) (*pb.PingPongResponse, error) {
	log.Println(in.GetPingMessage())
	return &pb.PingPongResponse{Result: "pong"}, nil
}
