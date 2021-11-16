package app

import (
	"log"
	"net/http"

	"github.com/evenlab/go-kit/context"
	"github.com/evenlab/go-kit/errors"
	"github.com/goava/di"
	"github.com/urfave/cli/v2"

	"keep_connection_server/config"
	api "keep_connection_server/errors"
	"keep_connection_server/rpc_server"
	"keep_connection_server/service"
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
	// invoke pong service starter
	if err := dic.Invoke(service.Pong.Start); err != nil {
		return api.ErrStartPongService(err)
	}

	// invoke rpc server starter
	if err := dic.Invoke(rpc_server.Start); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return api.ErrStartRPCServer(err)
		}
	}

	return nil
}

// provideServices tries to provide required
// services into application container.
func provideServices(dic *di.Container) error {
	// provide the app rpc server
	if err := dic.Provide(rpc_server.RPCServer); err != nil {
		return err
	}

	// provide the app keep connect service
	if err := dic.Provide(service.NewKeepConnectService); err != nil {
		return err
	}

	return nil
}
