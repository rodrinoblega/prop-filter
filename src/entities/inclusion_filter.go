package entities

type InclusionFilter struct {
	Field string
	Value bool
}

func (inf *InclusionFilter) Matches(property Property) bool {
	return property.HasAmenity(inf.Field) == inf.Value
}
