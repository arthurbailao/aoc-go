package day01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `1721
979
366
299
675
1456
`

func TestParse(t *testing.T) {
	expected := []int{1721, 979, 366, 299, 675, 1456}

	actual := parse(strings.NewReader(input))
	assert.Equal(t, expected, actual, "they should be equal")
}
