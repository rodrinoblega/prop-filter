package filters_provider

import (
	"github.com/rodrinoblega/prop-filter/src/entities"
)

type MockFilterProvider struct {
	filters *entities.Filters
}

func NewMockFilterProvider(filters *entities.Filters) *MockFilterProvider {
	return &MockFilterProvider{filters: filters}
}

func (fp *MockFilterProvider) GetFilters() *entities.Filters {
	return fp.filters
}
