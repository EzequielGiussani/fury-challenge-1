package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/items"
	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItemHandler_GetItemPriceSuccess(t *testing.T) {
	//Given
	iServiceMock := items.NewServiceMock()
	iServiceMock.HandleGetItem = func(ctx context.Context, itemID string) (*items.Item, apierrors.ApiError) {
		return &items.Item{
			ID:         "MLA-12345",
			Price:      2,
			CurrencyID: "peso ARG",
		}, nil
	}
	i := &ItemHandler{
		itemService: iServiceMock,
	}
	responseBodyExpected := `{"id":"MLA-12345","price":2,"currency_id":"peso ARG"}`
	//init app test server
	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	app.Router.Get("/prices/{id}", i.GetItemPrice)

	//When
	r := httptest.NewRequest(http.MethodGet, "/prices/MLA-12345", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)

	//then
	assert.EqualValues(t, http.StatusOK, w.Result().StatusCode)
	assert.EqualValues(t, responseBodyExpected, w.Body.String())
}

func TestItemHandler_GetItemPriceWhenFailingGetItem(t *testing.T) {
	//Given
	iServiceMock := items.NewServiceMock()
	iServiceMock.HandleGetItem = func(ctx context.Context, itemID string) (*items.Item, apierrors.ApiError) {
		return &items.Item{}, apierrors.NewInternalServerApiError("error getting item", nil)
	}
	i := &ItemHandler{
		itemService: iServiceMock,
	}
	responseBodyExpected := `{"message":"error getting item","error":"internal_server_error","status":500,"cause":[]}`
	//init app test server
	app, err := fury.NewWebApplication()
	require.NoError(t, err)
	app.Router.Get("/prices/{id}", i.GetItemPrice)

	//When
	r := httptest.NewRequest(http.MethodGet, "/prices/MLA-12345", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)

	//then
	assert.EqualValues(t, http.StatusInternalServerError, w.Result().StatusCode)
	assert.EqualValues(t, responseBodyExpected, w.Body.String())
}
