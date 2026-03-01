package service

import (
	"github.com/harshtripathi/postgres-with-go/models"
	"github.com/harshtripathi/postgres-with-go/repository"
)

// CreateStock handles business logic
func CreateStock(stock models.Stock) (int64, error) {
	return repository.InsertStock(stock)
}

// FetchAllStocks handles logic
func FetchAllStocks() ([]models.Stock, error) {
	return repository.GetAllStocks()
}

// FetchStock handles logic
func FetchStock(id int64) (models.Stock, error) {
	return repository.GetStockByID(id)
}
