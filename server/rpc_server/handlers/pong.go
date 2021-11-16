package handlers

import (
	"io"
	"log"
	"time"

	"github.com/jkrus/keep_connection/pb"
)

type Pong struct {
	pb.UnimplementedPingPongServer
	inspect   bool
	idleTimer *time.Timer
}

// PingMessage implements method PingPongServer.PingMessage.
func (p *Pong) PingMessage(stream pb.PingPong_PingMessageServer) error {
	if !p.inspect {
		p.inspect = true
		p.idleTimer = time.NewTimer(20 * time.Second)
		go p.watchTimer()
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if in.GetPingMessage() == "ping" {
			p.idleTimer.Reset(20 * time.Second)
			log.Println("client in touch")
			if err = stream.Send(&pb.PingPongResponse{Result: "pong"}); err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func (p *Pong) watchTimer() {
	defer func() {
		p.idleTimer.Stop()
	}()
	for {
		select {
		case <-p.idleTimer.C:
			log.Println("client is out of touch")
			p.inspect = false
			break
		}
	}
}
