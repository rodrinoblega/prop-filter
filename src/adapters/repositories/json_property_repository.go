package repositories

import (
	"encoding/json"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"os"
)

type JSONPropertyRepository struct {
	Properties []entities.Property
}

func NewJSONPropertyRepository(filename string) (*JSONPropertyRepository, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var properties []entities.Property
	if err := json.Unmarshal(file, &properties); err != nil {
		return nil, err
	}

	return &JSONPropertyRepository{Properties: properties}, nil
}

func (repo *JSONPropertyRepository) FindProperties() ([]entities.Property, error) {
	return repo.Properties, nil
}
