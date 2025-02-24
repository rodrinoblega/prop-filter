package filters_provider_go

import (
	"flag"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Args_Filters_Provider_Get_Filters(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cmd", "-minSqFt", "1100", "-maxSqFt", "1200", "-amenities", "garage:true", "-contains", "Family", "-lat", "34", "-lon", "-118", "-maxDist", "100"}

	filterProvider := NewArgsFilterProvider()

	filters, err := filterProvider.GetFilters()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var expectedFilters []entities.Filter
	expectedFilters = append(expectedFilters,
		&entities.SquareFootageFilter{SquareFootageRange: &entities.SquareFootageRange{Min: toPtr(1100), Max: toPtr(1200)}},
		&entities.InclusionFilter{Field: "garage", Value: true},
		&entities.MatchingFilter{Word: "Family"},
		&entities.DistanceFilter{Lat: 34, Lon: -118, MaxDist: 100})

	assert.Equal(t, len(expectedFilters), len(filters.Filters))
}

func TestParseSquareFootage_InvalidMin(t *testing.T) {
	flags := map[string]string{
		"minSqFt": "abc",
		"maxSqFt": "1234",
	}

	filters := parseSquareFootage(flags)

	if len(filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(filters))
	}

	sqFtFilter, ok := filters[0].(*entities.SquareFootageFilter)
	if !ok {
		t.Fatalf("Expected SquareFootageFilter, got %T", filters[0])
	}

	if sqFtFilter.SquareFootageRange.Min != nil {
		t.Errorf("Expected Min to be nil, got %d", *sqFtFilter.SquareFootageRange.Min)
	}
}

func TestParseSquareFootage_InvalidMax(t *testing.T) {
	flags := map[string]string{
		"minSqFt": "123",
		"maxSqFt": "cde",
	}

	filters := parseSquareFootage(flags)

	if len(filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(filters))
	}

	sqFtFilter, ok := filters[0].(*entities.SquareFootageFilter)
	if !ok {
		t.Fatalf("Expected SquareFootageFilter, got %T", filters[0])
	}

	if sqFtFilter.SquareFootageRange.Max != nil {
		t.Errorf("Expected Max to be nil, got %d", *sqFtFilter.SquareFootageRange.Max)
	}
}

func TestParseAmenities_EmptyString(t *testing.T) {
	filters := parseAmenities("")

	if len(filters) != 0 {
		t.Errorf("Expected empty filter list, got %d filters", len(filters))
	}
}

func TestParseAmenities_InvalidFormat(t *testing.T) {
	filters := parseAmenities("garage:true,pool")

	if len(filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(filters))
	}

	if f, ok := filters[0].(*entities.InclusionFilter); !ok || f.Field != "garage" || f.Value != true {
		t.Errorf("Expected InclusionFilter for garage:true, got %+v", f)
	}
}

func TestParseDistance_MissingArgs(t *testing.T) {
	tests := []struct {
		name  string
		flags map[string]string
	}{
		{"Missing lat", map[string]string{"lon": "34.56", "maxDist": "10"}},
		{"Missing lon", map[string]string{"lat": "-58.45", "maxDist": "10"}},
		{"Missing maxDist", map[string]string{"lat": "-58.45", "lon": "34.56"}},
		{"All missing", map[string]string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseDistance(tt.flags)
			if len(got) != 0 {
				t.Errorf("Expected no filters, got %v", got)
			}
		})
	}
}

func TestParseDistance_InvalidValues(t *testing.T) {
	tests := []struct {
		name  string
		flags map[string]string
	}{
		{"Invalid lat", map[string]string{"lat": "abc", "lon": "34.56", "maxDist": "10"}},
		{"Invalid lon", map[string]string{"lat": "-58.45", "lon": "xyz", "maxDist": "10"}},
		{"Invalid maxDist", map[string]string{"lat": "-58.45", "lon": "34.56", "maxDist": "asdasd"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseDistance(tt.flags)
			if len(got) != 0 {
				t.Errorf("Expected no filters, got %v", got)
			}
		})
	}
}

func toPtr(number int) *int {
	return &number
}
