package database

import (
	"database/sql"
	"log"

	"github.com/harshtripathi/postgres-with-go/config"
	_ "github.com/lib/pq"
)

// DB is a global connection pool
var DB *sql.DB

// InitDB initializes DB once (not per request)
func InitDB() {

	connStr := config.GetEnv("POSTGRES_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB Open Error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
