package service

type (
	// Pong describes service interface.
	Pong interface {
		// Pong tries set pong to client.
		Pong() error

		// Start tries start service.
		Start() error

		// Stop tries stop service.
		Stop() error
	}
)
