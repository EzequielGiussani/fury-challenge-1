package exchanges

import (
	"context"
	"time"

	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

// type IService interface {
// 	GetExchangeRate(ctx context.Context, currencyFrom string, currencyTo string) (*Exchange, apierrors.ApiError)
// }

type ConvertionRatioAPIResponse struct {
	Rate       float64   `json:"rate"`
	ValidUntil time.Time `json:"valid_until"`
}

type APIClient interface {
	GetConvertionRatio(ctx context.Context, fromCurrencyID, toCurrencyID string) (*ConvertionRatioAPIResponse, apierrors.ApiError)
}
