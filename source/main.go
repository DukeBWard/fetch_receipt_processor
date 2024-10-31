package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// driver function
func main() {
	router := chi.NewRouter()

	router.Post("/receipts/process", postReceiptHandler)
	router.Get("/receipts/{id}/points", getPointsHandler)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
