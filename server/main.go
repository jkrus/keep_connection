package main

import (
	"log"

	"github.com/evenlab/go-kit/context"
	"github.com/goava/di"

	"keep_connection_server/app"
	"keep_connection_server/config"
)

func main() {
	// create the application container
	c, err := di.New(
		// provide the application
		di.Provide(app.NewApp),
		// provide default config constructor
		di.Provide(config.NewConfig),
		// provide the application's context
		di.Provide(context.NewContext),
	)
	if err != nil {
		log.Fatal(err)
	}

	// invoke application starter
	if err = c.Invoke(app.Start); err != nil {
		log.Fatal(err)
	}
}
