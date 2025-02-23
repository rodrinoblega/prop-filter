package entities

type Filter interface {
	Matches(property Property) bool
}

type Filters struct {
	Filters []Filter
}

func (c *Filters) ApplyFilters(property Property) bool {
	for _, filter := range c.Filters {
		if !filter.Matches(property) {
			return false
		}
	}

	return true
}
