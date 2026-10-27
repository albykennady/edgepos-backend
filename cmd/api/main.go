package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albykennady/edgepos-backend/internal/httpapi"
)

const port = 8080

func main() {
	handler := httpapi.NewRouter()

	// TODO 1: Log that EdgePOS is starting on port 8080.
	address := fmt.Sprintf(":%d", port)
	log.Println("Starting application on port", address)

	// TODO 2: Start the HTTP server on port 8080 using handler.

	// TODO 3: Handle the error returned if the server fails to start.
	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatal("HTTP Server Failed ", err)
	}

}
