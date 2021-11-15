// Copyright © 2017-2020 The EVEN Foundation developers

package rpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

var (
	con *grpc.ClientConn // global variables are harmful
	ctx context.Context  // global variables are harmful

	cancel  func()
	timeout = 30 * time.Second
)

func connect() error {
	var err error
	// Set up a connection to the server
	con, err = grpc.Dial("localhost:4141", grpc.WithInsecure())

	if err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}

	freshContext()

	return nil
}

func terminate() {
	defer cancel()
	if con != nil {
		_ = con.Close()
	}
}

func freshContext() {
	// Create a context
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
}
