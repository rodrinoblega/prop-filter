package entities

import "strings"

type MatchingFilter struct {
	Word string
}

func (mf *MatchingFilter) Matches(property Property) bool {
	if strings.Contains(property.Description, mf.Word) {
		return true
	}

	return false
}
