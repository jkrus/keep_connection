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
	ConnIdle  time.Duration
}

// PingMessage implements method PingPongServer.PingMessage.
func (p *Pong) PingMessage(stream pb.PingPong_PingMessageServer) error {
	if !p.inspect {
		p.inspect = true
		p.idleTimer = time.NewTimer(p.ConnIdle)
		log.Println(p.ConnIdle)
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
			p.idleTimer.Reset(p.ConnIdle)
			log.Println("client in touch")
			//	md := metadata.Pairs("PING","1")
			log.Println("start delay")
			time.Sleep(5 * time.Second)
			log.Println("end delay")
			/*	if err = stream.SendHeader(md); err != nil {
				log.Println("Transport Error = ", err)
			}*/
			if err = stream.Send(&pb.PingPongResponse{Result: "pong"}); err != nil {
				log.Println("Err = ", err)
				return err
			}
			log.Println("sended")
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
