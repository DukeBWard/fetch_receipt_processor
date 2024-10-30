package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// just keep it in memory for now
var receipts = make(map[string]Receipt)

// POST /receipts/process
func postReceiptHandler(w http.ResponseWriter, r *http.Request) {

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
