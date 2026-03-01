package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harshtripathi/postgres-with-go/routers"
)

func main() {
	r := routers.Routers()

	fmt.Println("starting server on the port 8080 ...")

	log.Fatal(http.ListenAndServe(":8080", r))

}
