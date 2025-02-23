package filters_provider_go

import (
	"errors"
	"github.com/rodrinoblega/prop-filter/src/entities"
)

type MockFilterProvider struct {
	filters *entities.Filters
}

func NewMockFilterProvider(filters *entities.Filters) *MockFilterProvider {
	return &MockFilterProvider{filters: filters}
}

func (fp *MockFilterProvider) GetFilters() (*entities.Filters, error) {
	return fp.filters, nil
}

type MockErrorFilterProvider struct{}

func NewErrorMockFilterProvider() *MockErrorFilterProvider {
	return &MockErrorFilterProvider{}
}

func (fp *MockErrorFilterProvider) GetFilters() (*entities.Filters, error) {
	return nil, errors.New("mocked Error")
}
