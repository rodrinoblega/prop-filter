package filters_provider_go

import (
	"flag"
	"github.com/rodrinoblega/prop-filter/src/entities"
	"strconv"
	"strings"
)

type ArgsFilterProvider struct{}

type Args struct {
	Flags map[string]string
}

func NewArgsFilterProvider() *ArgsFilterProvider {
	return &ArgsFilterProvider{}
}

func (fp *ArgsFilterProvider) GetFilters() (*entities.Filters, error) {
	args := ParseFlags()

	var filters []entities.Filter

	filters = append(filters, parseSquareFootage(args.Flags)...)
	filters = append(filters, parseAmenities(args.Flags["amenities"])...)

	return &entities.Filters{Filters: filters}, nil
}

func ParseFlags() (args Args) {
	flags := make(map[string]string)

	minSqFt := flag.String("minSqFt", "", "Minimum square footage")
	maxSqFt := flag.String("maxSqFt", "", "Maximum square footage")
	amenities := flag.String("amenities", "", "Comma-separated list of amenities with true/false (e.g., garage:true,pool:false)")

	flag.Parse()

	if *minSqFt != "" {
		flags["minSqFt"] = *minSqFt
	}
	if *maxSqFt != "" {
		flags["maxSqFt"] = *maxSqFt
	}
	if *amenities != "" {
		flags["amenities"] = *amenities
	}

	return Args{Flags: flags}
}

func parseSquareFootage(flags map[string]string) []entities.Filter {
	var filters []entities.Filter

	if val, ok := flags["minSqFt"]; ok || flags["maxSqFt"] != "" {
		var minSqFt, maxSqFt *int
		if ok {
			v, err := parseInt(val)
			if err == nil {
				minSqFt = &v
			}
		}
		if val, ok := flags["maxSqFt"]; ok {
			v, err := parseInt(val)
			if err == nil {
				maxSqFt = &v
			}
		}

		filters = append(filters, &entities.SquareFootageFilter{
			SquareFootageRange: &entities.SquareFootageRange{Min: minSqFt, Max: maxSqFt},
		})
	}

	return filters
}

func parseAmenities(input string) []entities.Filter {
	var filters []entities.Filter

	if input == "" {
		return filters
	}

	pairs := strings.Split(input, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		if len(parts) != 2 {
			continue
		}

		amenity := parts[0]
		value := parts[1] == "true"
		filters = append(filters, &entities.InclusionFilter{Field: amenity, Value: value})
	}

	return filters
}

func parseInt(input string) (int, error) {
	v, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return v, nil
}
