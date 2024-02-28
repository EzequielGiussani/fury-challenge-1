package items

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients"
	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
	"github.com/mercadolibre/fury_go-meli-toolkit-restful/rest"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://test.mercadolibre.com"

const itemsGetResponse = `
	{
		"price": 12.23,
		"id": "%s",
		"currency_id": "ARS"
	}`

const responseBodyNotFoundError = `
	{ 
		"message": "Resource not found.",
		"error": "not_found", 
		"status": 404, 
		"cause": []
	}`

const responseBodyInternalError = `
	{ 
		"message": "Error getting item by id.",
		"error": "internal_error", 
		"status": 500, 
		"cause": []
	}`

func getMockServerConfig(statusCode int, path string, method string, response string, expectedCallCount int) *rest.Mock {
	mockServerConfig := new(rest.Mock)
	mockServerConfig.URL = baseURL + path
	mockServerConfig.HTTPMethod = method
	mockServerConfig.RespHTTPCode = statusCode
	mockServerConfig.RespBody = response
	mockServerConfig.ExpectedCallCount = expectedCallCount
	return mockServerConfig
}

func TestAPIClient(t *testing.T) {
	apiClient := NewClient(clients.Config{BaseURL: baseURL})
	client := apiClient.(Client)
	assert.Equal(t, baseURL, client.restClient.BaseURL)
}

func TestClientGetOK(t *testing.T) {
	defer rest.StopMockupServer()
	ctx := context.Background()

	uri, _ := url.Parse("/items/MLA2?attributes=price,id,currency_id")
	mockServerConfig := getMockServerConfig(http.StatusOK, uri.String(), http.MethodGet, fmt.Sprintf(itemsGetResponse, "MLA2"), 1)
	rest.StartMockupServer()
	err := rest.AddMockups(mockServerConfig)
	assert.Nil(t, err)

	client := NewClient(clients.Config{BaseURL: baseURL})

	itemDTO, err := client.GetItem(ctx, "MLA2")
	assert.Nil(t, err)

	assert.Equal(t, "MLA2", itemDTO.ItemID)
	assert.Equal(t, "ARS", itemDTO.CurrencyID)
	assert.Equal(t, 12.23, itemDTO.Price)
}

func TestClientGetNotFound(t *testing.T) {
	defer rest.StopMockupServer()
	ctx := context.Background()

	uri, _ := url.Parse("/items/MLA2?attributes=price,id,currency_id")
	mockServerConfig := getMockServerConfig(http.StatusNotFound, uri.String(), http.MethodGet, responseBodyNotFoundError, 1)
	rest.StartMockupServer()
	err := rest.AddMockups(mockServerConfig)
	assert.Nil(t, err)

	client := NewClient(clients.Config{BaseURL: baseURL})

	_, err = client.GetItem(ctx, "MLA2")
	assert.NotNil(t, err)
	assert.Equal(t, 404, err.(apierrors.ApiError).Status())
}

func TestClientGetInternalServerError(t *testing.T) {
	defer rest.StopMockupServer()
	ctx := context.Background()

	uri, _ := url.Parse("/items/MLA2?attributes=price,id,currency_id")
	mockServerConfig := getMockServerConfig(http.StatusInternalServerError, uri.String(), http.MethodGet, responseBodyInternalError, 1)
	rest.StartMockupServer()
	err := rest.AddMockups(mockServerConfig)
	assert.Nil(t, err)

	client := NewClient(clients.Config{BaseURL: baseURL})

	_, err = client.GetItem(ctx, "MLA2")
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.(apierrors.ApiError).Status())
}
