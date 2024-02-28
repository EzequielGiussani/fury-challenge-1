package handlers

import (
	"net/http"

	"github.com/mercadolibre/fury_bootcamp-go-demo/cmd/api/dependencies"
	"github.com/mercadolibre/fury_bootcamp-go-demo/internal/clients/items"
	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type ItemHandler struct {
	itemService items.IService
}

func NewItemHandler(depend dependencies.Dependencies) *ItemHandler {
	return &ItemHandler{
		itemService: depend.ItemsService,
	}
}

func (i *ItemHandler) GetItemPrice(w http.ResponseWriter, r *http.Request) error {
	params := web.Params(r)

	itemID, err := params.String("id")
	if err != nil {
		return web.EncodeJSON(w, "error fetching item id from url", http.StatusInternalServerError)
	}

	item, apiErr := i.itemService.GetItem(r.Context(), itemID)
	if apiErr != nil {
		return web.EncodeJSON(w, apiErr, apiErr.Status())
	}

	return web.EncodeJSON(w, item, http.StatusOK)
}

func (i *ItemHandler) GetItemPriceUSD(w http.ResponseWriter, r *http.Request) error {
	params := web.Params(r)

	itemID, err := params.String("id")
	if err != nil {
		return web.EncodeJSON(w, "error fetching item id from url", http.StatusInternalServerError)
	}

	item, apiErr := i.itemService.GetItemPriceUSD(r.Context(), itemID)
	if apiErr != nil {
		return web.EncodeJSON(w, apiErr, apiErr.Status())
	}

	return web.EncodeJSON(w, item, http.StatusOK)
}
