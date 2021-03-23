package day01_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day01"
	"github.com/stretchr/testify/assert"
)

const input = `1721
979
366
299
675
1456
`

var solution = day01.New(strings.NewReader(input))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 514579, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecond()

	if assert.NoError(t, err) {
		assert.Equal(t, 241861950, actual, "they should be equal")
	}
}
