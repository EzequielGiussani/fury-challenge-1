package handlers

import (
	"github.com/mercadolibre/fury_go-core/pkg/web"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, "pong", http.StatusOK)
}
