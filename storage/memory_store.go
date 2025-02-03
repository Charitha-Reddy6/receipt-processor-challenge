package storage

import (
	"receipt-processor/models"
	"sync"
)

var receiptStore = make(map[string]models.ProcessedReceipt)
var mu sync.Mutex

func SaveReceipt(id string, receipt models.ProcessedReceipt) {
	mu.Lock()
	defer mu.Unlock()
	receiptStore[id] = receipt
}

func GetReceipt(id string) (models.ProcessedReceipt, bool) {
	mu.Lock()
	defer mu.Unlock()
	receipt, exists := receiptStore[id]
	return receipt, exists
}
