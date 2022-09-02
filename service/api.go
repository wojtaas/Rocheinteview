package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"net/http"
	"rocheinteview"
)

type Service interface {
	Ping(message string) (rocheinteview.PingResponse, error)
}

type API struct {
	name    string
	service Service
}

func NewApi(name string, service Service) *API {
	return &API{
		name:    name,
		service: service,
	}
}

// GetHandler returns a router for HTTP API.
// @title Roche-Interview
// @version 1.0.0
// @BasePath /api
func (api *API) GetHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route(fmt.Sprintf("/api/v1/%s/", api.name), func(r chi.Router) {
		r.Post("/echo", api.EchoPing)
	})

	return r
}

// EchoPing
// @summary EchoPing
// @description Returns echo message from Postman, with additionally fields.
// @tags Roche-Interview
// @param message path string true "message"
// @failure 500 "server error"
// @failure 400 "not found field message in path"
// @success 200 {object} rocheinteview.PingResponse
// @Router /v1/roche/echo [post]
func (api *API) EchoPing(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has(rocheinteview.PingMessage) {
		http.Error(w, "not found message in path", http.StatusBadRequest)
		return
	}

	message := r.URL.Query().Get(rocheinteview.PingMessage)

	pingResponse, err := api.service.Ping(message)
	if err != nil {
		http.Error(w, fmt.Sprintf("api.service.Ping() %s", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(pingResponse); err != nil {
		http.Error(w, fmt.Sprintf("json.NewEncoder().Encode(): %s", err), http.StatusInternalServerError)
		return
	}
}
