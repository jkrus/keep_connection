package service

type (
	// Pong describes service interface.
	Pong interface {
		// Start tries start service.
		Start() error

		// Stop tries stop service.
		Stop() error
	}
)
