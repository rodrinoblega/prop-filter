package entities

import (
	"math"
)

type DistanceFilter struct {
	Lat     float64
	Lon     float64
	MaxDist float64
}

func (df *DistanceFilter) Matches(property Property) bool {
	const R = 6371

	lat1, lon1, lat2, lon2 := toRadians(df.Lat), toRadians(df.Lon), toRadians(property.Location[0]), toRadians(property.Location[1])

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	currentDistance := R * c

	if df.MaxDist < currentDistance {
		return false
	}

	return true
}

func toRadians(deg float64) float64 {
	return deg * (math.Pi / 180)
}
