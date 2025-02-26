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
	flags := map[string]string{"input": filename}

	reader, _ := NewJSONPropertyReader(flags)
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

func TestNewJSONPropertyReader_FileOpenError(t *testing.T) {
	_, err := NewJSONPropertyReader(map[string]string{"input": "invalid"})

	assert.Error(t, err)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open file: open invalid: no such file or directory")
}

func TestFindProperties_JSONDecodeError(t *testing.T) {
	invalidJSON := `[ {  
		"squareFootage"
	}]`

	filePath := createTempFileWithContent(t, invalidJSON)
	defer os.Remove(filePath)

	flags := map[string]string{"input": filePath}

	propertiesChan := make(chan entities.Property, 10)
	errorChan := make(chan error, 10)

	reader, _ := NewJSONPropertyReader(flags)
	go reader.FindProperties(propertiesChan, errorChan)

	err := <-errorChan
	if err == nil || !strings.Contains(err.Error(), "error decoding property: invalid character '}' after object key") {
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

	flags := map[string]string{"input": filePath}

	propertiesChan := make(chan entities.Property, 10)
	errorChan := make(chan error, 10)

	reader, _ := NewJSONPropertyReader(flags)
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
	flags := map[string]string{"input": filePath}

	propertiesChan := make(chan entities.Property, 10)
	errorChan := make(chan error, 10)

	reader, _ := NewJSONPropertyReader(flags)
	go reader.FindProperties(propertiesChan, errorChan)

	var lastErr error
	for err := range errorChan {
		lastErr = err
	}

	assert.Error(t, lastErr)
	assert.Contains(t, lastErr.Error(), "error reading end of JSON array")
}

func TestFindProperties_No_Input(t *testing.T) {
	flags := map[string]string{}

	_, err := NewJSONPropertyReader(flags)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required flag: --input (path to JSON file)")
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
