package config

import (
	"github.com/mercadolibre/fury_bootcamp-go-demo/cmd/api/dependencies"
	"github.com/mercadolibre/fury_bootcamp-go-demo/cmd/api/handlers"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func MapRoutes(app *fury.Application, depend dependencies.Dependencies) {
	app.Post("/ping", handlers.PingHandler)

	itemHandler := handlers.NewItemHandler(depend)
	app.Get("/prices/{id}", itemHandler.GetItemPrice)
	app.Get("/prices_usd/{id}", itemHandler.GetItemPriceUSD)
}
