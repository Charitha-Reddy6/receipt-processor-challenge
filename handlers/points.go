package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/storage"

	"github.com/gorilla/mux"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, exists := storage.GetReceipt(id)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": receipt.Points}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
