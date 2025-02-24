package entities

import "testing"

func TestDistanceFilter_Matches(t *testing.T) {
	tests := []struct {
		name     string
		filter   DistanceFilter
		property Property
		matches  bool
	}{
		{
			name: "Property outside max distance",
			filter: DistanceFilter{
				Lat:     40,
				Lon:     -74,
				MaxDist: 50,
			},
			property: Property{
				Location: [2]float64{34, -118},
			},
			matches: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.filter.Matches(tt.property)
			if got != tt.matches {
				t.Errorf("Expected %v, got %v", tt.matches, got)
			}
		})
	}
}
