package app

import (
	"log"
	"os"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"
	"github.com/urfave/cli/v2"

	"keep_connection_server/config"
)

// NewApp returns an application.
func NewApp(ctx context.Context, cfg *config.Config, dic *di.Container) *cli.App {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"ver"},
		Usage:   "print the version",
	}

	// construct cli application
	app := &cli.App{
		Name:    config.AppServerName,
		Usage:   config.AppUsage,
		Version: version(),
		ExitErrHandler: func(_ *cli.Context, err error) {
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	// register commands into cli application
	startCommand(ctx, cfg, dic, app)

	return app
}

// Start starts the application.
func Start(ctx context.Context, app *cli.App) error {
	return app.RunContext(ctx, os.Args)
}
