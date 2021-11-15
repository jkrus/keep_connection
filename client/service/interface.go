package service

type (
	// Ping describes service interface.
	Ping interface {
		// Ping tries set ping signal to server.
		Ping() error

		// Start tries start service.
		Start() error

		// Stop tries stop service.
		Stop() error
	}
)
