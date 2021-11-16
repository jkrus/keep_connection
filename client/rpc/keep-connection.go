package rpc

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/jkrus/keep_connection/pb"
)

// RenewalRequest is sent to confirm that the TCP/IP connection is still valid
func RenewalRequest() (bool, error) {
	defer terminate()
	if err := connect(); err != nil {
		return false, fmt.Errorf("rpc connect: %w", err)
	}
	client := pb.NewPingPongClient(con)
	stream, err := client.PingMessage(context.Background())
	if err != nil {
		return false, err
	}
	waitc := make(chan bool)
	waitPing := make(chan bool)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("keep connection failed to receive a pong : %v", err)
			}
			if in.GetResult() == "pong" {
				log.Println(in.GetResult())
				waitPing <- true
				return
			}
		}
	}()

	if err = stream.Send(&pb.PingPongRequest{PingMessage: "ping"}); err != nil {
		log.Fatalf("keep connection failed to send a ping: %v", err)
	}

	stream.CloseSend()
	<-waitPing
	return true, nil
}
