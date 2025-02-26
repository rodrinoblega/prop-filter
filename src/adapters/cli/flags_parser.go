package cli

import "flag"

func ParseFlags() map[string]string {
	flags := make(map[string]string)

	minSqFt := flag.String("minSqFt", "", "Minimum square footage")
	maxSqFt := flag.String("maxSqFt", "", "Maximum square footage")
	amenities := flag.String("amenities", "", "Comma-separated list of amenities with true/false (e.g., garage:true,pool:false)")
	contains := flag.String("contains", "", "Contains a specific string in the description")
	lat := flag.String("lat", "", "Latitude for location-based filtering")
	lon := flag.String("lon", "", "Longitude for location-based filtering")
	maxDist := flag.String("maxDist", "", "Maximum distance in kilometers")
	input := flag.String("input", "", "Name of JSON File")

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

	if *input != "" {
		flags["input"] = *input
	}

	return flags
}
