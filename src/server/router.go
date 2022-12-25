package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/config"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/handlers"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/repository/db"
	"github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/service"
)

var DBConfig config.DBConfig

func LoadRoute() {
	log.Println("INFO : Loading Router")
	router := mux.NewRouter().PathPrefix("/events-collector-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	log.Println("INFO : Router Loaded Successfully")
	log.Println("INFO : Application is started Successfully")
	http.ListenAndServe(":9999", router)
}

func registerAppRoutes(r *mux.Router) {
	log.Println("INFO : Registering Router ")

	log.Println("INFO : Registering Router ")
	eventRepo := db.NewDatabaseRepository(db.GetDatabaseProvider(DBConfig.DBUser, DBConfig.DBName, DBConfig.DBPassword))
	eventService := service.NewEventService(eventRepo)
	eventHandlers := handlers.NewHandler(eventService)

	r.HandleFunc("/collector/events", eventHandlers.HandleGetAllEvents).Methods(http.MethodGet)
	r.HandleFunc("/collector/events", eventHandlers.HandleGetAllEvents).Methods(http.MethodPost)
	log.Println("INFO : Router Registered Successfully")
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func LoadDB() {

	DBConfig.DBUser = os.Getenv("POSTGRES_USER")

	DBConfig.DBPassword = os.Getenv("POSTGRES_PASSWORD")

	DBConfig.DBName = os.Getenv("POSTGRES_DB")

}
