package day03_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day03"
	"github.com/stretchr/testify/assert"
)

const testCase = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

var solution = day03.New(strings.NewReader(testCase))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 7, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecond()

	if assert.NoError(t, err) {
		assert.Equal(t, 336, actual, "they should be equal")
	}
}
