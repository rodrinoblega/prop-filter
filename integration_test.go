package main

import (
	"errors"
	"github.com/rodrinoblega/prop-filter/src/adapters/filters_provider.go"
	"github.com/rodrinoblega/prop-filter/src/adapters/repositories"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/rodrinoblega/prop-filter/src/use_cases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompleteFlow_With_Test_Dependencies(t *testing.T) {
	repo := repositories.NewMockPropertyRepository(mockedProperties())

	filterProvider := filters_provider_go.NewMockFilterProvider(mockedFilters())

	propertyFinder := use_cases.NewPropertyFinder(repo, filterProvider)

	filteredProperties, err := propertyFinder.Execute()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedCount := 1
	if len(filteredProperties) != expectedCount {
		t.Errorf("Expected %d properties, got %d", expectedCount, len(filteredProperties))
	}
}

func TestCompleteFlow_With_Test_Dependencies_Filter_Provider_Error(t *testing.T) {
	repo := repositories.NewErrorMockPropertyRepository()

	filterProvider := filters_provider_go.NewMockFilterProvider(mockedFilters())

	propertyFinder := use_cases.NewPropertyFinder(repo, filterProvider)

	_, err := propertyFinder.Execute()
	assert.Error(t, err, errors.New("mocked error"))
}

func TestCompleteFlow_With_Test_Dependencies_Args_Provider_Error(t *testing.T) {
	repo := repositories.NewMockPropertyRepository(mockedProperties())

	filterProvider := filters_provider_go.NewErrorMockFilterProvider()

	propertyFinder := use_cases.NewPropertyFinder(repo, filterProvider)

	_, err := propertyFinder.Execute()
	assert.Error(t, err, errors.New("mocked error"))
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
