package handlers

import (
	"github.com/jkrus/keep_connection/pb"
)

type Pong struct {
	pb.UnimplementedPingPongServer
}

// PingMessage implements method PingPongServer.PingMessage.
func (p *Pong) PingMessage(stream pb.) error {
	stream.
	return  nil
}
