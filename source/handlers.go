package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// just keep it in memory for now
var receipts = make(map[string]Receipt)

// POST /receipts/process
func postReceiptHandler(w http.ResponseWriter, r *http.Request) {

	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate a unique id for the receipt
	id := uuid.New().String()
	receipt.ID = id
	receipts[id] = receipt

	// TODO: calculate points
	receipt.Points = calcPoints(receipt)

	res := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// GET /receipts/{id}/points
func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	receipt, err := receipts[id]
	if !err {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	res := map[string]int{
		"points": receipt.Points,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
