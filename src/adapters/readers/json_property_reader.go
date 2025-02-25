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
		sendError(errorChan, fmt.Errorf("failed to open file: %w", err))
		return
	}

	decoder := json.NewDecoder(file)

	if err := readJSONStart(decoder, errorChan); err != nil {
		return
	}

	readProperties(decoder, propertiesChan, errorChan)

	readJSONEnd(decoder, errorChan)

	close(propertiesChan)
	close(errorChan)
}

func readJSONStart(decoder *json.Decoder, errorChan chan error) error {
	if _, err := decoder.Token(); err != nil {
		sendError(errorChan, fmt.Errorf("error reading start of JSON array: %w", err))
		return err
	}
	return nil
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
