package main

import (
	"log"
	"net/http"
	"receipt-processor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", r)
}
