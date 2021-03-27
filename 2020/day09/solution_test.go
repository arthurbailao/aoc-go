package day09_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day09"
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

var solution = day09.New(strings.NewReader(input))

func TestSolveFirstWithParam(t *testing.T) {
	actual, err := solution.SolveFirstWithParam(5)

	if assert.NoError(t, err) {
		assert.Equal(t, 127, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecondWithParam(5)

	if assert.NoError(t, err) {
		assert.Equal(t, 62, actual, "they should be equal")
	}
}
