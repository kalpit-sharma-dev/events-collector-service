package handlers

import (
	"net/http"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/service"
)

type Handler struct {
	service service.EventCollectorService
}

func (h *Handler) HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleCreateEvents(w http.ResponseWriter, r *http.Request) {

}

func NewHandler(service service.EventCollectorService) *Handler {
	return &Handler{
		service: service,
	}
}
