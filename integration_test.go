package main

import (
	"github.com/rodrinoblega/prop-filter/src/adapters/filters_provider"
	"github.com/rodrinoblega/prop-filter/src/adapters/readers"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/rodrinoblega/prop-filter/src/use_cases"
	"testing"
)

func TestCompleteFlow_With_Test_Dependencies(t *testing.T) {
	propertyReader := readers.NewMockPropertyReader(mockedProperties())

	filterProvider := filters_provider.NewMockFilterProvider(mockedFilters())

	propertyFinder := use_cases.NewPropertyFinder(propertyReader, filterProvider)

	filteredProperties := propertyFinder.Execute()

	expectedCount := 1
	if len(filteredProperties) != expectedCount {
		t.Errorf("Expected %d properties, got %d", expectedCount, len(filteredProperties))
	}
}

func mockedFilters() *entities.Filters {
	minSqFt := 150
	maxSqFt := 200

	filters := &entities.Filters{
		Filters: []entities.Filter{
			&entities.SquareFootageFilter{
				SquareFootageRange: &entities.SquareFootageRange{
					Min: &minSqFt,
					Max: &maxSqFt,
				},
			},
			&entities.InclusionFilter{
				Field: "garage",
				Value: true,
			},
			&entities.MatchingFilter{
				Word: "Family",
			},
			&entities.DistanceFilter{
				Lat:     34,
				Lon:     -118,
				MaxDist: 100,
			},
		},
	}

	return filters
}

func mockedProperties() []entities.Property {
	properties := []entities.Property{
		{
			SquareFootage: 120,
			Amenities: map[string]bool{
				"garage": true,
			},
			Description: "Beautiful apartment.",
			Location:    [2]float64{40, -110},
		},
		{
			SquareFootage: 200,
			Amenities: map[string]bool{
				"garage": true,
			},
			Description: "Family home.",
			Location:    [2]float64{34, -118},
		},
		{
			SquareFootage: 80,
			Description:   "Cozy house.",
			Location:      [2]float64{41, -87},
		},
	}

	return properties
}
