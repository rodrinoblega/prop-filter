package entities

import (
	"testing"
)

func Test_SquareFootageFilter_Matches(t *testing.T) {
	tests := []struct {
		name     string
		filter   *SquareFootageFilter
		property Property
		expected bool
	}{
		{"Property within range", &SquareFootageFilter{&SquareFootageRange{Min: toPtr(100), Max: toPtr(200)}}, Property{SquareFootage: 150}, true},
		{"Property below range", &SquareFootageFilter{&SquareFootageRange{Min: toPtr(100), Max: toPtr(200)}}, Property{SquareFootage: 50}, false},
		{"Property above range", &SquareFootageFilter{&SquareFootageRange{Min: toPtr(100), Max: toPtr(200)}}, Property{SquareFootage: 250}, false},
		{"No min, within max", &SquareFootageFilter{&SquareFootageRange{Min: nil, Max: toPtr(200)}}, Property{SquareFootage: 150}, true},
		{"No max, above min", &SquareFootageFilter{&SquareFootageRange{Min: toPtr(100), Max: nil}}, Property{SquareFootage: 150}, true},
		{"Nil min, nil max", &SquareFootageFilter{&SquareFootageRange{Min: nil, Max: nil}}, Property{SquareFootage: 500}, true},
		{"Filter is nil", &SquareFootageFilter{nil}, Property{SquareFootage: 500}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.filter.Matches(tt.property)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func toPtr(i int) *int {
	return &i
}
