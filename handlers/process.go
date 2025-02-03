package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"

	"github.com/google/uuid"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := utils.CalculatePoints(receipt)

	processedReceipt := models.ProcessedReceipt{
		ID:     id,
		Points: points,
	}

	storage.SaveReceipt(id, processedReceipt)
	response := map[string]string{"id": id}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
