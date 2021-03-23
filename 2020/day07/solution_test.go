package day07_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day07"
	"github.com/stretchr/testify/assert"
)

const input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

var solution = day07.New(strings.NewReader(input))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 4, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecond()

	if assert.NoError(t, err) {
		assert.Equal(t, 32, actual, "they should be equal")
	}
}
