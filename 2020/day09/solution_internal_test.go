package day09

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestParse(t *testing.T) {

	expected := []int{35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576}

	actual := parse(strings.NewReader(input))
	assert.Equal(t, expected, actual, "they should be equal")
}
