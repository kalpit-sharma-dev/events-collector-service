package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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
	var req models.Request
	err := decoder.Decode(&req)
	if err != nil {
		log.Fatal("unable to decode", err)
		return
	}
	log.Println("INFO : HandleCreateUrl", r)

	url, err := h.service.CreateEvent(r.Context(), req)
	if err != nil {
		log.Fatal("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		//errors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}

func NewHandler(service service.EventCollectorService) *Handler {
	return &Handler{
		service: service,
	}
}
