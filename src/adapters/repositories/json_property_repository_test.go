package repositories

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/stretchr/testify/assert"
)

func Test_NewJSONPropertyRepository_Find_Properties(t *testing.T) {
	expectedProperties := []entities.Property{
		{SquareFootage: 150},
		{SquareFootage: 200},
	}

	filename := createTempJSONFile(t, expectedProperties)
	defer os.Remove(filename)

	repo, err := NewJSONPropertyRepository(filename)
	properties, err := repo.FindProperties()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.NoError(t, err)
	assert.Len(t, properties, 2)
	assert.Equal(t, expectedProperties, properties)
}

func Test_NewJSONPropertyRepository_Find_Properties_FileNotFound(t *testing.T) {
	_, err := NewJSONPropertyRepository("no_exist.json")
	assert.Error(t, err)
}

func Test_NewJSONPropertyRepository_Find_Properties_FileNotFound_InvalidJSON(t *testing.T) {
	filename := createTempJSONFile(t, []entities.Property{})
	defer os.Remove(filename)

	os.WriteFile(filename, []byte("{invalid json}"), 0644)

	_, err := NewJSONPropertyRepository(filename)
	assert.Error(t, err)
}

func createTempJSONFile(t *testing.T, data []entities.Property) string {
	file, err := os.CreateTemp("", "test_json_repository.json")
	assert.NoError(t, err)

	jsonData, err := json.Marshal(data)
	assert.NoError(t, err)

	_, err = file.Write(jsonData)
	assert.NoError(t, err)

	file.Close()
	return file.Name()
}
