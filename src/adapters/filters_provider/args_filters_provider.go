package filters_provider

import (
	"github.com/rodrinoblega/prop-filter/src/entities"
	"strconv"
	"strings"
)

type ArgsFilterProvider struct {
	Args map[string]string
}

func NewArgsFilterProvider(flags map[string]string) *ArgsFilterProvider {
	return &ArgsFilterProvider{Args: flags}
}

func (fp *ArgsFilterProvider) GetFilters() *entities.Filters {
	var filters []entities.Filter

	filters = createFiltersBasedOnArgs(filters, fp.Args)

	return &entities.Filters{Filters: filters}
}

func createFiltersBasedOnArgs(filters []entities.Filter, args map[string]string) []entities.Filter {
	filters = append(filters, parseSquareFootage(args)...)
	filters = append(filters, parseAmenities(args["amenities"])...)
	filters = append(filters, parseContains(args["contains"])...)
	filters = append(filters, parseDistance(args)...)
	return filters
}

func parseSquareFootage(flags map[string]string) []entities.Filter {
	var filters []entities.Filter

	var minSqFt, maxSqFt *int

	if val, ok := flags["minSqFt"]; ok {
		if v, err := parseInt(val); err == nil {
			minSqFt = &v
		}
	}

	if val, ok := flags["maxSqFt"]; ok {
		if v, err := parseInt(val); err == nil {
			maxSqFt = &v
		}
	}

	if minSqFt != nil || maxSqFt != nil {
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

func parseContains(s string) []entities.Filter {
	var filters []entities.Filter

	if s != "" {
		filters = append(filters, &entities.MatchingFilter{Word: s})
	}

	return filters
}

func parseDistance(flags map[string]string) []entities.Filter {
	var filters []entities.Filter

	latStr, latOk := flags["lat"]
	lonStr, lonOk := flags["lon"]
	maxDistStr, distOk := flags["maxDist"]

	if !latOk || !lonOk || !distOk {
		return filters
	}

	lat, errLat := strconv.ParseFloat(latStr, 64)
	lon, errLon := strconv.ParseFloat(lonStr, 64)
	maxDist, errMaxDist := strconv.ParseFloat(maxDistStr, 64)

	if errLat != nil || errLon != nil || errMaxDist != nil {

		return filters
	}

	filters = append(filters, &entities.DistanceFilter{
		Lat:     lat,
		Lon:     lon,
		MaxDist: maxDist,
	})

	return filters
}

func parseInt(input string) (int, error) {
	v, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return v, nil
}
