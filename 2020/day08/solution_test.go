package day08_test

import (
	"strings"
	"testing"

	"github.com/arthurbailao/aoc/2020/day08"
	"github.com/stretchr/testify/assert"
)

const input = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

var solution = day08.New(strings.NewReader(input))

func TestSolveFirst(t *testing.T) {
	actual, err := solution.SolveFirst()

	if assert.NoError(t, err) {
		assert.Equal(t, 5, actual, "they should be equal")
	}
}

func TestSolveSecond(t *testing.T) {
	actual, err := solution.SolveSecond()

	if assert.NoError(t, err) {
		assert.Equal(t, 8, actual, "they should be equal")
	}
}
