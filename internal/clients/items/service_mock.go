package items

import (
	"context"

	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

type ServiceMock struct {
	HandleGetItem func(ctx context.Context, itemID string) (*Item, apierrors.ApiError)
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (sm ServiceMock) GetItem(ctx context.Context, itemID string) (*Item, apierrors.ApiError) {
	if sm.HandleGetItem != nil {
		return sm.HandleGetItem(ctx, itemID)
	}
	return &Item{}, nil
}

func (sm ServiceMock) GetItemPriceUSD(ctx context.Context, itemID string) (*Item, apierrors.ApiError) {
	return &Item{}, nil
}
