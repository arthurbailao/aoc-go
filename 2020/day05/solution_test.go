package day05_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day05"
	"github.com/stretchr/testify/assert"
)

const input = `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

var solution = day05.New(strings.NewReader(input))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 820, actual, "they should be equal")
	}
}
