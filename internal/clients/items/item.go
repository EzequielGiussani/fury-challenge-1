package items

import (
	"context"

	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

type Item struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	CurrencyID string  `json:"currency_id"`
}

type ItemAPIResponse struct {
	ItemID     string  `json:"id"`
	Price      float64 `json:"price"`
	CurrencyID string  `json:"currency_id"`
}

type IService interface {
	GetItem(ctx context.Context, itemID string) (*Item, apierrors.ApiError)
	GetItemPriceUSD(ctx context.Context, itemID string) (*Item, apierrors.ApiError)
}

type APIClient interface {
	GetItem(ctx context.Context, itemID string) (*ItemAPIResponse, apierrors.ApiError)
}
