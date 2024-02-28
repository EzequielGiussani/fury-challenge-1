package items

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients"
	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/apierrors"
	"github.com/mercadolibre/fury_go-meli-toolkit-goutils/logger"
	"github.com/mercadolibre/fury_go-meli-toolkit-restful/rest"
)

const (
	itemAttributes = "price,id,currency_id"
)

type Client struct {
	restClient *rest.RequestBuilder
}

func NewClient(config clients.Config) APIClient {
	customPool := &rest.CustomPool{
		MaxIdleConnsPerHost: 100,
	}
	headers := make(http.Header)

	restClient := &rest.RequestBuilder{
		BaseURL:        config.BaseURL,
		Headers:        headers,
		Timeout:        3 * time.Second,
		ContentType:    rest.JSON,
		EnableCache:    false,
		DisableTimeout: false,
		CustomPool:     customPool,
		MetricsConfig:  rest.MetricsReportConfig{TargetId: "items-api"},
	}
	return Client{restClient: restClient}
}

func (c Client) GetItem(ctx context.Context, itemID string) (*ItemAPIResponse, apierrors.ApiError) {
	var response *rest.Response
	query := clients.Query()
	query.Add("attributes", itemAttributes)

	uri, err := clients.BuildURL([]string{"/items", itemID}, query)
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("error parsing items URL", err)
	}
	response = c.restClient.Get(uri, rest.Context(ctx))

	if response.Err != nil || response.Response == nil {
		errMsg := "Unexpected error getting item"
		logger.Errorf(fmt.Sprintf("[item_id:%s] %s, url: %s", itemID, errMsg, uri), response.Err)
		return nil, apierrors.NewInternalServerApiError(errMsg, response.Err)
	}

	if response.StatusCode == http.StatusNotFound {
		return nil, apierrors.NewNotFoundApiError("item not found")
	}

	if response.StatusCode != http.StatusOK {
		errMsg := "Unexpected error getting item response"
		logger.Errorf(fmt.Sprintf("[item_id:%s] %s, url: %s, status code: %d", itemID, errMsg, uri, response.StatusCode), response.Err)
		return nil, apierrors.NewApiError(errMsg, "Unexpected error getting item", response.StatusCode, apierrors.CauseList{})
	}
	var itemAPIResponse ItemAPIResponse
	rawItem := response.Bytes()
	if unmarshalError := json.Unmarshal(rawItem, &itemAPIResponse); unmarshalError != nil {
		errMsg := "Unexpected error unmarshalling item"
		logger.Errorf(fmt.Sprintf("[item_id:%s] %s, value: %s", itemID, errMsg, string(rawItem)), unmarshalError)
		return nil, apierrors.NewInternalServerApiError(errMsg, unmarshalError)
	}

	return &itemAPIResponse, nil
}
