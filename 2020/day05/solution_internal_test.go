package day05

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

func TestParse(t *testing.T) {
	expected := []boardingPass{
		boardingPass{
			FrontBack: []string{"B", "F", "F", "F", "B", "B", "F"},
			RightLeft: []string{"R", "R", "R"},
		},
		boardingPass{
			FrontBack: []string{"F", "F", "F", "B", "B", "B", "F"},
			RightLeft: []string{"R", "R", "R"},
		},
		boardingPass{
			FrontBack: []string{"B", "B", "F", "F", "B", "B", "F"},
			RightLeft: []string{"R", "L", "L"},
		},
	}

	actual := parse(strings.NewReader(input))
	assert.Equal(t, expected, actual, "they should be equal")
}
