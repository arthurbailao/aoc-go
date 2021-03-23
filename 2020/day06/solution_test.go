package day06_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day06"
	"github.com/stretchr/testify/assert"
)

const input = `abc

a
b
c

ab
ac

a
a
a
a

b`

var solution = day06.New(strings.NewReader(input))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 11, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecond()

	if assert.NoError(t, err) {
		assert.Equal(t, 6, actual, "they should be equal")
	}
}
