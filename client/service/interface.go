package service

type (
	// KeepConnect describes service interface.
	KeepConnect interface {
		// Start tries start service.
		Start() error

		// Stop tries stop service.
		Stop() error
	}
)
