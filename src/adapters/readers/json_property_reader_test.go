package readers

import (
	"os"
	"strings"
	"testing"

	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/stretchr/testify/assert"
)

func Test_NewJSONPropertyRepository_Find_Properties(t *testing.T) {
	jsonData := `[  
		{"squareFootage": 150},  
		{"squareFootage": 200}  
	]`

	expectedProperties := []entities.Property{
		{SquareFootage: 150},
		{SquareFootage: 200},
	}

	filename := createTempFile(t, jsonData)
	defer os.Remove(filename)

	reader := NewJSONPropertyReader(filename)
	resultChan := make(chan entities.Property, 100)
	errorChan := make(chan error, 10)

	reader.FindProperties(resultChan, errorChan)

	var properties []entities.Property
	for prop := range resultChan {
		properties = append(properties, prop)
	}

	assert.Len(t, properties, 2)
	assert.Equal(t, expectedProperties, properties)
}

func TestFindProperties_FileOpenError(t *testing.T) {
	jobChan := make(chan entities.Property)
	errorChan := make(chan error)

	reader := JSONPropertyReader{filePath: "non_existent_file.json"}

	go reader.FindProperties(jobChan, errorChan)

	err := <-errorChan
	if err == nil || !strings.Contains(err.Error(), "failed to open file") {
		t.Errorf("Expected file open error, got: %v", err)
	}
}

func TestFindProperties_JSONDecodeError(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "invalid.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString("{invalid_json}")
	tmpFile.Close()

	jobChan := make(chan entities.Property)
	errorChan := make(chan error)

	reader := JSONPropertyReader{filePath: tmpFile.Name()}

	go reader.FindProperties(jobChan, errorChan)

	err = <-errorChan
	if err == nil || !strings.Contains(err.Error(), "error decoding property") {
		t.Errorf("Expected JSON decode error, got: %v", err)
	}
}

func TestFindProperties_ErrorMissingOpeningBracket(t *testing.T) {
	invalidJSON := `? {  
		"squareFootage": 1200,  
		"lighting": "high",  
		"price": 250000  
	}]`

	filePath := createTempFileWithContent(t, invalidJSON)
	defer os.Remove(filePath)

	propertiesChan := make(chan entities.Property, 10)
	errorChan := make(chan error, 10)

	reader := NewJSONPropertyReader(filePath)
	go reader.FindProperties(propertiesChan, errorChan)

	err := <-errorChan
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error reading start of JSON array: invalid character '?' looking for beginning of value")
}

func TestFindProperties_ErrorMissingClosingBracket(t *testing.T) {
	invalidJSON := `[  
	{  
		"squareFootage": 1200,  
		"lighting": "high",  
		"price": 250000  
	}`

	filePath := createTempFileWithContent(t, invalidJSON)
	defer os.Remove(filePath)

	propertiesChan := make(chan entities.Property, 10)
	errorChan := make(chan error, 10)

	reader := NewJSONPropertyReader(filePath)
	go reader.FindProperties(propertiesChan, errorChan)

	var lastErr error
	for err := range errorChan {
		lastErr = err
	}

	assert.Error(t, lastErr)
	assert.Contains(t, lastErr.Error(), "error reading end of JSON array")
}

func createTempFileWithContent(t *testing.T, content string) string {
	file, err := os.CreateTemp("", "test_invalid_json.json")
	assert.NoError(t, err)

	_, err = file.WriteString(content)
	assert.NoError(t, err)

	file.Close()
	return file.Name()
}

func createTempFile(t *testing.T, content string) string {
	file, err := os.CreateTemp("", "test_json.json")
	assert.NoError(t, err)

	_, err = file.WriteString(content)
	assert.NoError(t, err)

	file.Close()
	return file.Name()
}
