package readers

import (
	"github.com/rodrinoblega/prop-filter/src/entities"
)

type MockPropertyReader struct {
	properties []entities.Property
}

func NewMockPropertyReader(properties []entities.Property) *MockPropertyReader {
	return &MockPropertyReader{properties: properties}
}

func (m *MockPropertyReader) FindProperties(result chan<- entities.Property, errors chan error) {
	for _, property := range m.properties {
		result <- property
	}
	close(result)
}
