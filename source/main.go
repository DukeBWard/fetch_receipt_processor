package main

import (
	"log"

	"github.com/go-chi/chi"
)

// driver function
func main() {
	router := chi.NewRouter()

	router.Post("/receipts/process", postReceiptHandler)
	router.Get("/receipts/{id}/points", getPointsHandler)

	log.Println("Server is running on port 8080")
}
