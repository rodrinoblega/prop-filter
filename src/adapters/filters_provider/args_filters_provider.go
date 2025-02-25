package filters_provider

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

func (fp *ArgsFilterProvider) GetFilters() *entities.Filters {
	args := ParseFlags()

	var filters []entities.Filter

	filters = createFiltersBasedOnArgs(filters, args)

	return &entities.Filters{Filters: filters}
}

func createFiltersBasedOnArgs(filters []entities.Filter, args Args) []entities.Filter {
	filters = append(filters, parseSquareFootage(args.Flags)...)
	filters = append(filters, parseAmenities(args.Flags["amenities"])...)
	filters = append(filters, parseContains(args.Flags["contains"])...)
	filters = append(filters, parseDistance(args.Flags)...)
	return filters
}

func ParseFlags() (args Args) {
	flags := make(map[string]string)

	minSqFt := flag.String("minSqFt", "", "Minimum square footage")
	maxSqFt := flag.String("maxSqFt", "", "Maximum square footage")
	amenities := flag.String("amenities", "", "Comma-separated list of amenities with true/false (e.g., garage:true,pool:false)")
	contains := flag.String("contains", "", "Contains a specific string in the description")
	lat := flag.String("lat", "", "Latitude for location-based filtering")
	lon := flag.String("lon", "", "Longitude for location-based filtering")
	maxDist := flag.String("maxDist", "", "Maximum distance in kilometers")

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

	if *contains != "" {
		flags["contains"] = *contains
	}

	if *lat != "" {
		flags["lat"] = *lat
	}

	if *lon != "" {
		flags["lon"] = *lon
	}

	if *maxDist != "" {
		flags["maxDist"] = *maxDist
	}

	return Args{Flags: flags}
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
