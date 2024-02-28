package items

import (
	"context"

	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
)

type ClientMock struct {
	HandleGetItem func(ctx context.Context, itemID string) (*ItemAPIResponse, apierrors.ApiError)
}

func NewClientMock() *ClientMock {
	return &ClientMock{}
}

func (m ClientMock) GetItem(ctx context.Context, itemID string) (*ItemAPIResponse, apierrors.ApiError) {
	if m.HandleGetItem != nil {
		return m.HandleGetItem(ctx, itemID)
	}
	return nil, nil
}
