package service

type (
	// Ping describes service interface.
	Ping interface {
		// Start tries start service.
		Start() error

		// Stop tries stop service.
		Stop() error
	}
)
