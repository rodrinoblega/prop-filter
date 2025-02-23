package entities

import "testing"

type MockFilter struct {
	Result bool
}

func (m *MockFilter) Matches(_ Property) bool {
	return m.Result
}

func Test_Filters_ApplyFilters(t *testing.T) {
	tests := []struct {
		name     string
		filters  []Filter
		expected bool
	}{
		{"All filters pass", []Filter{&MockFilter{true}, &MockFilter{true}}, true},
		{"One filter fails", []Filter{&MockFilter{true}, &MockFilter{false}}, false},
		{"All filters fail", []Filter{&MockFilter{false}, &MockFilter{false}}, false},
		{"No filters (should pass)", []Filter{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filters := Filters{Filters: tt.filters}
			result := filters.ApplyFilters(Property{})
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
