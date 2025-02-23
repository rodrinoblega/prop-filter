package entities

type Property struct {
	SquareFootage int             `json:"squareFootage"`
	Lighting      string          `json:"lighting"`
	Price         float64         `json:"price"`
	Rooms         int             `json:"rooms"`
	Bathrooms     int             `json:"bathrooms"`
	Location      [2]float64      `json:"location"`
	Description   string          `json:"description"`
	Amenities     map[string]bool `json:"amenities"`
}

type SquareFootageRange struct {
	Min *int
	Max *int
}

func (sfr *SquareFootageRange) Contains(value int) bool {
	if sfr.Min != nil && value < *sfr.Min {
		return false
	}
	if sfr.Max != nil && value > *sfr.Max {
		return false
	}

	return true
}

func (p *Property) HasAmenity(amenity string) bool {
	return p.Amenities[amenity]
}
