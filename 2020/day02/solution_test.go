package day02_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day02"
	"github.com/stretchr/testify/assert"
)

const testCase = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

var solution = day02.New(strings.NewReader(testCase))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 2, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecond()

	if assert.NoError(t, err) {
		assert.Equal(t, 1, actual, "they should be equal")
	}
}
