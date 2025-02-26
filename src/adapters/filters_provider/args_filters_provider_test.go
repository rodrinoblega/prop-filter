package filters_provider

import (
	"github.com/rodrinoblega/prop-filter/src/entities"
	"testing"
)

func TestParseContains_ValidCase(t *testing.T) {
	flags := map[string]string{
		"contains": "match",
	}

	filterProvider := NewArgsFilterProvider(flags)
	filters := filterProvider.GetFilters().Filters

	if len(filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(filters))
	}

	matchingFilter, ok := filters[0].(*entities.MatchingFilter)
	if !ok {
		t.Fatalf("Expected MatchingFilter, got %T", filters[0])
	}

	if matchingFilter.Word != "match" {
		t.Errorf("Expected Word to be match, got %s", matchingFilter.Word)
	}
}

func TestParseSquareFootage_InvalidMin(t *testing.T) {
	flags := map[string]string{
		"minSqFt": "abc",
		"maxSqFt": "1234",
	}

	filterProvider := NewArgsFilterProvider(flags)
	filters := filterProvider.GetFilters().Filters

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

	filterProvider := NewArgsFilterProvider(flags)
	filters := filterProvider.GetFilters().Filters

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

func TestParseDistance_ValidValue(t *testing.T) {
	flags := map[string]string{
		"lat":     "34",
		"lon":     "34",
		"maxDist": "34",
	}

	filterProvider := NewArgsFilterProvider(flags)
	filters := filterProvider.GetFilters().Filters

	if len(filters) != 1 {
		t.Errorf("Expected 1 filter, got %d", len(filters))
	}

	distanceFilter, ok := filters[0].(*entities.DistanceFilter)
	if !ok {
		t.Fatalf("Expected MatchingFilter, got %T", filters[0])
	}

	if distanceFilter.Lat != 34 {
		t.Errorf("Expected Lat to be 34, got %f", distanceFilter.Lat)
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
