// Copyright © 2017-2020 The EVEN Foundation developers

package rpc

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc/keepalive"

	"evenfound/even/server/api"
	"evenfound/even/util/prompt"

	"evenfound/evenctl/config"

	"google.golang.org/grpc"
)

const (
	deadline = 60 * time.Second
)

var (
	con *grpc.ClientConn // global variables are harmful
	ctx context.Context  // global variables are harmful

	cancel  func()
	timeout = 30 * time.Second // why not constant?
)

func connect() error {
	var err error
	// Set up a connection to the server
	con, err = grpc.Dial(config.RPCAddress, grpc.WithInsecure())
	grpc.KeepaliveParams(keepalive.ClientParameters{
		PermitWithoutStream: true,
		Time:                10,
		Timeout:             5,
	})
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

func preparePasswords(rq *api.CreateRequest) error {
	var err error
	if rq.Password.Phrase == "" {
		// Start by prompting for the private passphrase.
		rq.Password.Phrase, err = prompt.EnterPassword()
		if err != nil {
			return err
		}
	}

	return err
}

func prepareSeedWithMnemonic(rq *api.CreateRequest) error {
	var err error
	reader := bufio.NewReader(os.Stdin)
	if rq.Seed.Phrase == "" {
		// Ascertain the wallet generation mnemonic.
		rq.Mnemonic.Phrase, err = prompt.Mnemonic(reader, rq.Mnemonic.Phrase, rq.Password.Phrase)
		if err != nil {
			return err
		}
		// Ascertain the wallet generation seed.
		if rq.Mnemonic.Phrase == "" {
			rq.Seed.Phrase, err = prompt.Seed(reader, false)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
