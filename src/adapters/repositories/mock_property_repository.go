package repositories

import (
	"errors"
	"github.com/rodrinoblega/prop-filter/src/entities"
)

type MockPropertyRepository struct {
	properties []entities.Property
}

func NewMockPropertyRepository(properties []entities.Property) *MockPropertyRepository {
	return &MockPropertyRepository{properties: properties}
}

func (m *MockPropertyRepository) FindProperties() ([]entities.Property, error) {
	return m.properties, nil
}

type ErrorMockPropertyRepository struct{}

func NewErrorMockPropertyRepository() *ErrorMockPropertyRepository {
	return &ErrorMockPropertyRepository{}
}

func (m *ErrorMockPropertyRepository) FindProperties() ([]entities.Property, error) {
	return nil, errors.New("mocked error")
}
