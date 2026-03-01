package middlewares

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harshtripathi/postgres-with-go/models"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

// ------------------- CREATE -------------------

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	insertID := insertStock(stock)

	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

// ------------------- GET ONE -------------------

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	stock, err := getStock(int64(id))
	if err != nil {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(stock)
}

func getStock(id int64) (models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stock models.Stock

	sqlStatement := `SELECT stockid, name, price, company FROM stocks WHERE stockid=$1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	if err != nil {
		return stock, err
	}

	return stock, nil
}

// ------------------- GET ALL -------------------

func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStock()
	if err != nil {
		http.Error(w, "Error fetching stocks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stocks)
}

func getAllStock() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stocks []models.Stock

	rows, err := db.Query(`SELECT stockid, name, price, company FROM stocks`)
	if err != nil {
		return stocks, err
	}
	defer rows.Close()

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
		if err != nil {
			return stocks, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

// ------------------- UPDATE -------------------

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	updatedRows := updateStock(int64(id), stock)

	msg := fmt.Sprintf("Stock updated. Rows affected: %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func updateStock(id int64, stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)
	if err != nil {
		panic(err)
	}

	rowsAffected, _ := res.RowsAffected()
	return rowsAffected
}

// ------------------- DELETE -------------------

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	deletedRows := deleteStock(int64(id))

	msg := fmt.Sprintf("Stock deleted. Rows affected: %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func deleteStock(id int64) int64 {
	db := createConnection()
	defer db.Close()

	res, err := db.Exec(`DELETE FROM stocks WHERE stockid=$1`, id)
	if err != nil {
		panic(err)
	}

	rowsAffected, _ := res.RowsAffected()
	return rowsAffected
}
