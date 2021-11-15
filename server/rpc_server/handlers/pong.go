package handlers

import (
	"keep_connection/pb"
)

type Pong struct {
	pb.UnimplementedPingPongServer
}

// PingMessage implements method PingPongServer.PingMessage.
func (p *Pong) PingMessage(stream pb.PingPong_PingMessageServer) error {
	stream.
	return  nil
}
