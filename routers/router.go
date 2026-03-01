package routers

import (
	"github.com/gorilla/mux"
	"github.com/harshtripathi/postgres-with-go/middlewares"
)

func Routers() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/stock", middlewares.CreateStock).Methods("POST")
	router.HandleFunc("/api/stock", middlewares.GetAllStock).Methods("GET")
	router.HandleFunc("/api/stock/{id}", middlewares.GetStock).Methods("GET")
	router.HandleFunc("/api/stock/{id}", middlewares.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stock/{id}", middlewares.DeleteStock).Methods("DELETE")

	return router
}
