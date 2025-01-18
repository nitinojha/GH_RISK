package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitinojha/GH_RISK/models"
	"github.com/nitinojha/GH_RISK/storage"
)

// GetAllRisks handles GET /v1/risks
func GetAllRisks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	risks := storage.GetRisks()
	json.NewEncoder(w).Encode(risks)
}

// CreateRisk handles POST /v1/risks
func CreateRisk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var risk models.Risk
	if err := json.NewDecoder(r.Body).Decode(&risk); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newRisk, err := storage.AddRisk(risk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newRisk)
}

// GetRiskByID handles GET /v1/risks/{id}
func GetRiskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	riskID := vars["id"]
	risk, err := storage.GetRiskByID(riskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(risk)
}
