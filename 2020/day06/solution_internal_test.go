package day06

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `abc

a
b
c

ab
ac`

func TestParse(t *testing.T) {
	expected := []group{
		group{
			Answers:         map[string]int{"a": 1, "b": 1, "c": 1},
			NumOfPassengers: 1,
		},
		group{
			Answers:         map[string]int{"a": 1, "b": 1, "c": 1},
			NumOfPassengers: 3,
		},
		group{
			Answers:         map[string]int{"a": 2, "b": 1, "c": 1},
			NumOfPassengers: 2,
		},
	}

	actual := parse(strings.NewReader(input))
	assert.Equal(t, expected, actual, "they should be equal")
}
