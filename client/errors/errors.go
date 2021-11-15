package errors

import (
	"github.com/evenlab/go-kit/errors"
)

const (
	errProvideCliContextMsg = "provide cli context"
	errStartPingServiceMsg  = "start ping service"
)

func ErrProvideCliContext(w error) error {
	return errors.WrapErr(errProvideCliContextMsg, w)
}

func ErrStartPingService(w error) error {
	return errors.WrapErr(errStartPingServiceMsg, w)
}
