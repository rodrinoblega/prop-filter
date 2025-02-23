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
	os.Args = []string{"cmd", "-minSqFt", "1100", "-maxSqFt", "1200"}

	filterProvider := NewArgsFilterProvider()

	filters, err := filterProvider.GetFilters()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var expectedFilters []entities.Filter
	expectedFilters = append(expectedFilters,
		&entities.SquareFootageFilter{SquareFootageRange: &entities.SquareFootageRange{Min: toPtr(1100), Max: toPtr(1200)}})

	assert.Equal(t, len(expectedFilters), len(filters.Filters))
}

func toPtr(number int) *int {
	return &number
}
