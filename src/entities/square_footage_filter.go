package entities

type SquareFootageFilter struct {
	SquareFootageRange *SquareFootageRange
}

type SquareFootageRange struct {
	Min *int
	Max *int
}

func (sff *SquareFootageFilter) Matches(property Property) bool {
	if sff.SquareFootageRange == nil {
		return true
	}

	return sff.SquareFootageRange.Contains(property.SquareFootage)
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
