package items

import (
	"context"
	"net/http"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/exchanges"
	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

const (
	USDCurrency = "USD"
)

var (
	CurrencyRates = make(map[string]RateValidUntil)
)

type RateValidUntil struct {
	Rate       float64
	ValidUntil time.Time
}

type Service struct {
	itemsClient     APIClient
	exchangesClient exchanges.APIClient
}

func NewService(itemsClient APIClient, exchangesClient exchanges.APIClient) *Service {
	return &Service{
		itemsClient:     itemsClient,
		exchangesClient: exchangesClient,
	}
}

func (s *Service) GetItem(ctx context.Context, itemID string) (*Item, apierrors.ApiError) {
	itemResponse, err := s.itemsClient.GetItem(ctx, itemID)
	if err != nil || itemResponse == nil {
		return nil, err
	}
	return &Item{ID: itemResponse.ItemID, Price: itemResponse.Price, CurrencyID: itemResponse.CurrencyID}, err
}

func (s *Service) GetItemPriceUSD(ctx context.Context, itemID string) (*Item, apierrors.ApiError) {
	itemResponse, err := s.itemsClient.GetItem(ctx, itemID)

	if err != nil || itemResponse == nil {
		return nil, err
	}

	if value, ok := CurrencyRates[itemResponse.CurrencyID]; ok {

		if time.Now().UTC().After(value.ValidUntil) {
			err := storeRatio(s, ctx, itemResponse, USDCurrency)

			if err != nil || itemResponse == nil {
				return nil, apierrors.NewApiError(err.Error(), err.Error(), http.StatusInternalServerError, apierrors.CauseList{})
			}
		}

	} else {
		err := storeRatio(s, ctx, itemResponse, USDCurrency)

		if err != nil || itemResponse == nil {
			return nil, apierrors.NewApiError(err.Error(), err.Error(), http.StatusInternalServerError, apierrors.CauseList{})
		}
	}

	return &Item{ID: itemResponse.ItemID, Price: getConvertedPrice(itemResponse.Price, CurrencyRates[itemResponse.CurrencyID].Rate), CurrencyID: USDCurrency}, err

}

func storeRatio(s *Service, ctx context.Context, itemResponse *ItemAPIResponse, toCurrency string) (err error) {
	convResponse, err := s.exchangesClient.GetConvertionRatio(ctx, itemResponse.CurrencyID, toCurrency)

	if err != nil || itemResponse == nil {
		return
	}

	CurrencyRates[itemResponse.CurrencyID] = RateValidUntil{
		Rate:       convResponse.Rate,
		ValidUntil: convResponse.ValidUntil,
	}

	return
}

func getConvertedPrice(currentPrice float64, rate float64) float64 {
	return currentPrice * rate
}
