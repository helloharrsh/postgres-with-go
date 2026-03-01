package routers

import (
	"github.com/gorilla/mux"
	"github.com/harshtripathi/postgres-with-go/handlers"
)

// SetupRouter registers routes
func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/stock", handlers.CreateStock).Methods("POST")
	router.HandleFunc("/api/stock", handlers.GetAllStocks).Methods("GET")
	router.HandleFunc("/api/stock/{id}", handlers.GetStockByID).Methods("GET")

	return router
}
