package filters_provider_go

import (
	"flag"
	"github.com/rodrinoblega/prop-filter/src/entities"
)

type ArgsFilterProvider struct{}

func NewArgsFilterProvider() *ArgsFilterProvider {
	return &ArgsFilterProvider{}
}

func (fp *ArgsFilterProvider) GetFilters() (*entities.Filters, error) {
	minSqFt, maxSqFt := ParseFlags()

	flag.Parse()

	var filters []entities.Filter
	if minSqFt != nil || maxSqFt != nil {
		filters = append(filters,
			&entities.SquareFootageFilter{SquareFootageRange: &entities.SquareFootageRange{Min: minSqFt, Max: maxSqFt}})
	}

	return &entities.Filters{Filters: filters}, nil
}

func ParseFlags() (minSqFt, maxSqFt *int) {
	minSqFt = flag.Int("minSqFt", 0, "Minimum square footage")
	maxSqFt = flag.Int("maxSqFt", 0, "Maximum square footage")
	flag.Parse()
	return
}
