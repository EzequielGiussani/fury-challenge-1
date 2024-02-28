package main

import (
	"log"
	"os"

	"github.com/mercadolibre/fury_bootcamp-go-demo/cmd/api/config"
	"github.com/mercadolibre/fury_bootcamp-go-demo/cmd/api/dependencies"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/platform/environment"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	env := environment.GetFromString(os.Getenv("GO_ENVIRONMENT"))

	depend, err := dependencies.BuildDependencies(env)
	if err != nil {
		return err
	}

	config.MapRoutes(app, depend)

	return app.Run()
}
