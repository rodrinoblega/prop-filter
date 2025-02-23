package entities

type SquareFootageFilter struct {
	SquareFootageRange *SquareFootageRange
}

func (sff *SquareFootageFilter) Matches(property Property) bool {
	if sff.SquareFootageRange == nil {
		return true
	}

	return sff.SquareFootageRange.Contains(property.SquareFootage)
}
