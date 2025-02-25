package readers

import (
	"encoding/json"
	"fmt"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"os"
)

type JSONPropertyReader struct {
	filePath string
}

func NewJSONPropertyReader(filePath string) *JSONPropertyReader {
	return &JSONPropertyReader{filePath: filePath}
}

func (r *JSONPropertyReader) FindProperties(propertiesChan chan<- entities.Property, errorChan chan error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		errorChan <- fmt.Errorf("failed to open file: %w", err)
		close(errorChan)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	//Reading first '['
	if _, err := decoder.Token(); err != nil {
		errorChan <- fmt.Errorf("error reading start of JSON array: %w", err)
		close(errorChan)
		return
	}

	for decoder.More() {
		var prop entities.Property
		if err := decoder.Decode(&prop); err != nil {
			errorChan <- fmt.Errorf("error decoding property: %w", err)
			continue
		}

		propertiesChan <- prop
	}

	//Reading last ']'
	if _, err := decoder.Token(); err != nil {
		errorChan <- fmt.Errorf("error reading end of JSON array: %w", err)
	}

	close(propertiesChan)
	close(errorChan)
}
