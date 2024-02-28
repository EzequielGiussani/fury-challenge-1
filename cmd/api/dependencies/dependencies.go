package dependencies

import (
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/exchanges"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/items"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/platform/environment"
	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

type Dependencies struct {
	ItemsService items.IService
}

func BuildDependencies(env environment.Environment) (Dependencies, apierrors.ApiError) {
	mlBaseURL := "https://api.mercadolibre.com"
	switch env {
	case environment.Production:
		mlBaseURL = "http://internal.mercadolibre.com"
	}
	itemsClient := items.NewClient(clients.Config{BaseURL: mlBaseURL})
	exchangesClient := exchanges.NewClient(clients.Config{BaseURL: mlBaseURL})
	itemsService := items.NewService(itemsClient, exchangesClient)

	return Dependencies{
		ItemsService: itemsService,
	}, nil
}
