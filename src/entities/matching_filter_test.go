package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MatchingFilter_Matches_No_Match(t *testing.T) {
	matchFilter := MatchingFilter{Word: "Familyyyy"}

	property := Property{Description: "Big living"}

	assert.False(t, matchFilter.Matches(property))
}
