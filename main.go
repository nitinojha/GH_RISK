package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitinojha/GH_RISK/handlers"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/v1/risks", handlers.GetAllRisks).Methods(http.MethodGet)
	r.HandleFunc("/v1/risks", handlers.CreateRisk).Methods(http.MethodPost)
	r.HandleFunc("/v1/risks/{id}", handlers.GetRiskByID).Methods(http.MethodGet)

	// Start the server
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
