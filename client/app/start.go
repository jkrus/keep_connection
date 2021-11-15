package app

import (
	"log"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"
	"github.com/urfave/cli/v2"

	"keep_connection_client/config"
	api "keep_connection_client/errors"
	"keep_connection_client/service"
)

// startCommand appends start action to application.
func startCommand(ctx context.Context, cfg *config.Config, dic *di.Container, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + config.AppUsage,
		Before: func(cc *cli.Context) error {
			// provide cli context
			if err := dic.Provide(func() *cli.Context { return cc }); err != nil {
				return api.ErrProvideCliContext(err)
			}
			// load application config
			if err := cfg.Init(); err != nil {
				return err
			}

			return provideServices(dic)
		},
		Action: func(cc *cli.Context) error {
			return invokeServices(dic)
		},
		After: func(cc *cli.Context) error {
			<-cc.Done() // wait while context canceled

			ctx.Cancel()
			ctx.WgWait() // wait while all workers finished

			log.Println("Application shutdown complete.")

			return nil
		},
	})
}

// invokeServices tries to invoke required
// services from application container.
func invokeServices(dic *di.Container) error {
	// invoke auth admin service starter
	if err := dic.Invoke(service.Ping.Start); err != nil {
		return api.ErrStartPingService(err)
	}

	return nil
}

// provideServices tries to provide required
// services into application container.
func provideServices(dic *di.Container) error {
	// provide the app admin service
	if err := dic.Provide(service.NewPingService); err != nil {
		return err
	}

	return nil
}
