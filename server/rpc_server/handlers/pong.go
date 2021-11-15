package handlers

import (
	"context"

	ping "evenfound/even/pingpong"
)

type Pong struct {
	api.UnsafePingPongServer
}

var _ api.PingPongServer = (*Pong)(nil)

// PingMessage implements method PingPongServer.PingMessage.
func (p *Pong) PingMessage(_ context.Context, in *api.PingPongRequest) (*api.PingPongResponse, error) {
	var reqpmsg = in.GetPingMessage()
	resmsg, err := getResponseMessage(reqpmsg)
	if err != nil {
		return nil, err
	}
	return &api.PingPongResponse{Result: resmsg}, nil
}

// getResponseMessage gets a string with the response "PONG"
func getResponseMessage(s string) (string, error) {
	return ping.GetMessage(s)
}
