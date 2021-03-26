package day08

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99`

func TestParse(t *testing.T) {
	expected := []instruction{
		instruction{Cmd: "nop", Val: 0},
		instruction{Cmd: "acc", Val: 1},
		instruction{Cmd: "jmp", Val: 4},
		instruction{Cmd: "acc", Val: 3},
		instruction{Cmd: "jmp", Val: -3},
		instruction{Cmd: "acc", Val: -99},
	}

	actual := parse(strings.NewReader(input))
	assert.Equal(t, expected, actual, "they should be equal")
}
