package storage

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/nitinojha/GH_RISK/models"
)

// MemoryStorage represents an in-memory storage for risks
type MemoryStorage struct {
	data map[string]models.Risk
	mu   sync.RWMutex
}

// Storage is the global in-memory storage instance
var Storage = MemoryStorage{
	data: make(map[string]models.Risk),
}

// GetRisks retrieves all risks
func GetRisks() []models.Risk {
	Storage.mu.RLock()
	defer Storage.mu.RUnlock()

	risks := make([]models.Risk, 0, len(Storage.data))
	for _, risk := range Storage.data {
		risks = append(risks, risk)
	}
	return risks
}

// AddRisk adds a new risk to the storage
func AddRisk(risk models.Risk) (models.Risk, error) {
	if err := risk.Validate(); err != nil {
		return models.Risk{}, err
	}

	Storage.mu.Lock()
	defer Storage.mu.Unlock()

	risk.ID = uuid.New().String()
	Storage.data[risk.ID] = risk
	return risk, nil
}

// GetRiskByID retrieves a risk by its ID
func GetRiskByID(id string) (models.Risk, error) {
	Storage.mu.RLock()
	defer Storage.mu.RUnlock()

	risk, exists := Storage.data[id]
	if !exists {
		return models.Risk{}, errors.New("risk not found")
	}
	return risk, nil
}
