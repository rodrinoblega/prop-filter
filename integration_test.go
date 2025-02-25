package main

import (
	"errors"
	"github.com/rodrinoblega/prop-filter/src/adapters/filters_provider"
	"github.com/rodrinoblega/prop-filter/src/adapters/readers"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/rodrinoblega/prop-filter/src/use_cases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompleteFlow_With_Test_Dependencies(t *testing.T) {
	reader := readers.NewMockPropertyReader(mockedProperties())

	filterProvider := filters_provider.NewMockFilterProvider(mockedFilters())

	propertyFinder := use_cases.NewPropertyFinder(
		use_cases.PropertyFinderInputs{
			PropertyReader: reader,
			FilterProvider: filterProvider},
	)

	filteredProperties, err := propertyFinder.Execute()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedCount := 1
	if len(filteredProperties) != expectedCount {
		t.Errorf("Expected %d properties, got %d", expectedCount, len(filteredProperties))
	}
}

func TestCompleteFlow_With_Test_Dependencies_Args_Provider_Error(t *testing.T) {
	reader := readers.NewMockPropertyReader(mockedProperties())

	filterProvider := filters_provider.NewErrorMockFilterProvider()

	propertyFinder := use_cases.NewPropertyFinder(
		use_cases.PropertyFinderInputs{
			PropertyReader: reader,
			FilterProvider: filterProvider},
	)

	_, err := propertyFinder.Execute()
	assert.Error(t, err, errors.New("mocked error"))
}

/*func TestCompleteFlow_With_Test_Dependencies_Error_In_ErrorChan(t *testing.T) {
	reader := readers.NewMockPropertyReader(mockedProperties())

	filterProvider := filters_provider.NewMockFilterProvider(mockedFilters())

	errorChan := make(chan error, 10)

	errorChan <- fmt.Errorf("failed to open file: %w", errors.New("custom error"))
	propertyFinder := use_cases.NewPropertyFinder(
		use_cases.PropertyFinderInputs{
			PropertyReader: reader,
			FilterProvider: filterProvider},
	)

	_, err := propertyFinder.Execute()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	select {
	case err := <-errorChan:
		if err != nil {
			t.Fatal("Expected a nil error because it should have been displayed within the execution")
		}
	default:
		//Empty error channel as expected
	}
}*/

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
