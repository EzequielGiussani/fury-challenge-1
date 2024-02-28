package exchanges

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
		MetricsConfig:  rest.MetricsReportConfig{TargetId: "exchanges-api"},
	}
	return Client{restClient: restClient}
}

type Client struct {
	restClient *rest.RequestBuilder
}

func (c Client) GetConvertionRatio(ctx context.Context, fromCurrencyID, toCurrencyID string) (*ConvertionRatioAPIResponse, apierrors.ApiError) {
	var response *rest.Response
	query := clients.Query()
	query.Add("from", fromCurrencyID)
	query.Add("to", toCurrencyID)

	uri, err := clients.BuildURL([]string{"/currency_conversions/search"}, query)

	if err != nil {
		return nil, apierrors.NewInternalServerApiError("error parsing currency_conversions URL", err)
	}

	response = c.restClient.Get(uri, rest.Context(ctx))

	if response.Err != nil || response.Response == nil {
		errMsg := "Unexpected error getting currency"
		logger.Errorf(fmt.Sprintf("[from_currency_id:%s][to_currency_id:%s] %s, url: %s", fromCurrencyID, toCurrencyID, errMsg, uri), response.Err)
		return nil, apierrors.NewInternalServerApiError(errMsg, response.Err)
	}

	if response.StatusCode == http.StatusNotFound {
		return nil, apierrors.NewNotFoundApiError("currency not found")
	}

	var convertionRatioAPIResponse ConvertionRatioAPIResponse
	rawItem := response.Bytes()
	if unmarshalError := json.Unmarshal(rawItem, &convertionRatioAPIResponse); unmarshalError != nil {
		errMsg := "Unexpected error unmarshalling item"
		logger.Errorf(fmt.Sprintf("[from_currency_id:%s][to_currency_id:%s], %s value: %s", fromCurrencyID, toCurrencyID, errMsg, string(rawItem)), unmarshalError)
		return nil, apierrors.NewInternalServerApiError(errMsg, unmarshalError)
	}

	return &convertionRatioAPIResponse, nil

}
