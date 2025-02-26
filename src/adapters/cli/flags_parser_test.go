package cli

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Parse_Flags(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cmd", "-minSqFt", "1100", "-maxSqFt", "1200", "-amenities", "garage:true", "-contains", "Family", "-lat", "34", "-lon", "-118", "-maxDist", "100", "-input", "properties.json"}

	flags := ParseFlags()

	expectedFlags := map[string]string{
		"minSqFt":   "1100",
		"maxSqFt":   "1200",
		"amenities": "garage:true",
		"contains":  "Family",
		"lat":       "34",
		"lon":       "-118",
		"maxDist":   "100",
		"input":     "properties.json",
	}

	assert.Equal(t, expectedFlags, flags)
}
