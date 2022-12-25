package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/handlers"
)

func LoadRoute() {
	log.Println("INFO : Loading Router")
	router := mux.NewRouter().PathPrefix("/events-collector-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	log.Println("INFO : Router Loaded Successfully")
	log.Println("INFO : Application is started Successfully")
	http.ListenAndServe(":8080", router)
}

func registerAppRoutes(r *mux.Router) {
	log.Println("INFO : Registering Router ")

	r.HandleFunc("/collector/events", handlers.HandleGetAllEvents).Methods(http.MethodGet)
	r.HandleFunc("/collector/events", handlers.HandleGetAllEvents).Methods(http.MethodPost)
	log.Println("INFO : Router Registered Successfully")
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
