package errors

import (
	"github.com/evenlab/go-kit/errors"
)

const (
	errProvideCliContextMsg = "provide cli context"
	errStartPongServiceMsg  = "start ping service"
	errStartRPCSererMsg     = "start rpc server"
)

func ErrProvideCliContext(w error) error {
	return errors.WrapErr(errProvideCliContextMsg, w)
}

func ErrStartPongService(w error) error {
	return errors.WrapErr(errStartPongServiceMsg, w)
}

func ErrStartRPCServer(w error) error {
	return errors.WrapErr(errStartRPCSererMsg, w)
}
