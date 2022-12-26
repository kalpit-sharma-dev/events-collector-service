package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/errors"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/service"
)

type Handler struct {
	service service.EventCollectorService
}

func (h *Handler) HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	resp, err := h.service.GetAllEvents(r.Context())
	if err != nil {
		log.Println("error", err)
		//errors.HandleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(resp)

}

func (h *Handler) HandleCreateEvents(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req models.Event
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("unable to decode", err)
		return
	}
	log.Println("INFO : HandleCreateEvent", r)

	event, err := h.service.CreateEvent(r.Context(), req)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		errors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func NewHandler(service service.EventCollectorService) *Handler {
	return &Handler{
		service: service,
	}
}
