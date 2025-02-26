package readers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"os"
)

type JSONPropertyReader struct {
	file *os.File
}

func NewJSONPropertyReader(flags map[string]string) (*JSONPropertyReader, error) {
	filePath, exists := flags["input"]
	if !exists || filePath == "" {
		return nil, errors.New("missing required flag: --input (path to JSON file)")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open file: %v", err))
	}

	return &JSONPropertyReader{file: file}, nil
}

func (r *JSONPropertyReader) FindProperties(propertiesChan chan<- entities.Property, errorChan chan error) {
	decoder := json.NewDecoder(r.file)

	readJSONStart(decoder, errorChan)

	readProperties(decoder, propertiesChan, errorChan)

	readJSONEnd(decoder, errorChan)

	close(propertiesChan)
	close(errorChan)
}

func readJSONStart(decoder *json.Decoder, errorChan chan error) {
	if _, err := decoder.Token(); err != nil {
		sendError(errorChan, fmt.Errorf("error reading start of JSON array: %w", err))
	}
}

func readProperties(decoder *json.Decoder, propertiesChan chan<- entities.Property, errorChan chan error) {
	for decoder.More() {
		var prop entities.Property
		if err := decoder.Decode(&prop); err != nil {
			sendError(errorChan, fmt.Errorf("error decoding property: %w", err))
			continue
		}
		propertiesChan <- prop
	}
}

func readJSONEnd(decoder *json.Decoder, errorChan chan error) {
	if _, err := decoder.Token(); err != nil {
		sendError(errorChan, fmt.Errorf("error reading end of JSON array: %w", err))
	}
}

func sendError(errorChan chan error, err error) {
	select {
	case errorChan <- err:
	default:
		// avoid blocking if channel is full
	}
}
