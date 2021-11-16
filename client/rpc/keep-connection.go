package rpc

import (
	"fmt"
	"log"

	"github.com/jkrus/keep_connection/pb"
)

// RenewalRequest is sent to confirm that the TCP/IP connection is still valid
func RenewalRequest(client pb.PingPongClient) (bool, error) {
	defer terminate()
	if err := connect(); err != nil {
		return false, fmt.Errorf("rpc connect: %w", err)
	}
client.
	request := &pb.PingPongRequest{PingMessage: "ping"}
	response, err := pb.NewPingPongClient(con).PingMessage(ctx, request)
	if err != nil {
		return false, fmt.Errorf("keep connection: %w", err)
	}

	log.Println(response.Result)

	return true, nil
}
