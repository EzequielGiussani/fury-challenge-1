package items

import (
	"context"
	"testing"

	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
	"github.com/stretchr/testify/assert"
)

func TestService_GetItemOK(t *testing.T) {
	//Given
	itemClientMock := NewClientMock()
	itemClientMock.HandleGetItem = func(ctx context.Context, itemID string) (*ItemAPIResponse, apierrors.ApiError) {
		return &ItemAPIResponse{
			ItemID:     "MLA-1",
			Price:      2.01,
			CurrencyID: "USD",
		}, nil
	}
	s := &Service{
		itemsClient: itemClientMock,
	}

	//Then
	itemObtained, err := s.GetItem(context.TODO(), "MLA-1")

	//When
	assert.Nil(t, err)
	assert.EqualValues(t, "MLA-1", itemObtained.ID)
	assert.EqualValues(t, 2.01, itemObtained.Price)
	assert.EqualValues(t, "USD", itemObtained.CurrencyID)
}

func TestService_GetItemWithErrorFromClient(t *testing.T) {
	//Given
	itemClientMock := NewClientMock()
	itemClientMock.HandleGetItem = func(ctx context.Context, itemID string) (*ItemAPIResponse, apierrors.ApiError) {
		return nil, apierrors.NewInternalServerApiError("mock error getting item", nil)
	}
	s := &Service{
		itemsClient: itemClientMock,
	}

	//then
	itemObtained, err := s.GetItem(context.TODO(), "MLA-1")

	//When
	assert.EqualValues(t, apierrors.NewInternalServerApiError("mock error getting item", nil), err)
	assert.Nil(t, itemObtained)
}
