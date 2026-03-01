package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harshtripathi/postgres-with-go/models"
	"github.com/harshtripathi/postgres-with-go/service"
)

// CreateStock API
func CreateStock(w http.ResponseWriter, r *http.Request) {

	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	id, err := service.CreateStock(stock)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// GetAllStocks API
func GetAllStocks(w http.ResponseWriter, r *http.Request) {

	stocks, err := service.FetchAllStocks()
	if err != nil {
		http.Error(w, "Error fetching stocks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stocks)
}

// GetStockByID API
func GetStockByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	stock, err := service.FetchStock(int64(id))
	if err != nil {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(stock)
}
