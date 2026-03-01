package main

import (
	"log"
	"net/http"

	"github.com/harshtripathi/postgres-with-go/config"
	"github.com/harshtripathi/postgres-with-go/database"
	"github.com/harshtripathi/postgres-with-go/routers"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	// Initialize DB once
	database.InitDB()

	router := routers.SetupRouter()

	log.Println("Server running on port 8080")

	http.ListenAndServe(":8080", router)
}
