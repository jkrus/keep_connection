package rpc

import (
	"evenfound/even/server/api"
	"fmt"
	"log"
)

// Ping is sent to confirm that the TCP/IP connection is still valid
func Ping(cmd string) error {
	defer terminate()
	if err := connect(); err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}

	request := &api.PingPongRequest{PingMessage: cmd}
	response, err := api.NewPingPongClient(con).PingMessage(ctx, request)
	if err != nil {
		return fmt.Errorf("ping-pong: %w", err)
	}

	log.Println(response.Result)

	return nil
}
