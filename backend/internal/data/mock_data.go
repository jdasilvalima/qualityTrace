package data

import (
    "encoding/json"
    "os"
    "sync"
    "backend/internal/models"
)

type MockDataStore struct {
    sync.RWMutex
    Users              []models.User              `json:"users"`
    Suppliers          []models.Supplier          `json:"suppliers"`
    Ingredients        []models.Ingredient        `json:"ingredients"`
    IngredientsReceived []models.IngredientReceived `json:"ingredientsReceived"`
    filePath          string
}

func NewMockDataStore(filePath string) (*MockDataStore, error) {
    store := &MockDataStore{
        filePath: filePath,
    }
    if err := store.loadData(); err != nil {
        return nil, err
    }
    return store, nil
}

func (s *MockDataStore) loadData() error {
    s.Lock()
    defer s.Unlock()

    data, err := os.ReadFile(s.filePath)
    if err != nil {
        return err
    }

    return json.Unmarshal(data, s)
}

// Ingredient methods
func (s *MockDataStore) GetIngredients() []models.Ingredient {
    s.RLock()
    defer s.RUnlock()
    return s.Ingredients
}