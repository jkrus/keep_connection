package rpc

import (
	"fmt"
	"strconv"

	"google.golang.org/grpc"

	"keep_connection_client/config"
)

var (
	con  *grpc.ClientConn // global variables are harmful
	conf *config.Config
)

func RPCClientConfig(cfg *config.Config) {
	conf = cfg
	return
}

func connect() error {
	var err error
	// Set up a connection to the server
	addr := conf.Host + ":" + strconv.Itoa(conf.Port)
	con, err = grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}

	return nil
}

func terminate() {
	if con != nil {
		_ = con.Close()
	}
}
