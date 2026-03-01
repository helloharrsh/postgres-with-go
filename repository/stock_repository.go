package repository

import (
	"github.com/harshtripathi/postgres-with-go/database"
	"github.com/harshtripathi/postgres-with-go/models"
)

// InsertStock inserts new stock into DB
func InsertStock(stock models.Stock) (int64, error) {

	var id int64

	err := database.DB.QueryRow(
		"INSERT INTO stocks(name, price, company) VALUES ($1,$2,$3) RETURNING stockid",
		stock.Name,
		stock.Price,
		stock.Company,
	).Scan(&id)

	return id, err
}

// GetAllStocks fetches all stocks
func GetAllStocks() ([]models.Stock, error) {

	rows, err := database.DB.Query("SELECT stockid,name,price,company FROM stocks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock

	for rows.Next() {
		var s models.Stock
		err := rows.Scan(&s.StockID, &s.Name, &s.Price, &s.Company)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}

	return stocks, nil
}

// GetStockByID fetches single stock
func GetStockByID(id int64) (models.Stock, error) {

	var s models.Stock

	err := database.DB.QueryRow(
		"SELECT stockid,name,price,company FROM stocks WHERE stockid=$1",
		id,
	).Scan(&s.StockID, &s.Name, &s.Price, &s.Company)

	return s, err
}
