package day02

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testCase = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestParse(t *testing.T) {
	expected := []policyAndPassword{
		policyAndPassword{policy{[]int{1, 3}, "a"}, "abcde"},
		policyAndPassword{policy{[]int{1, 3}, "b"}, "cdefg"},
		policyAndPassword{policy{[]int{2, 9}, "c"}, "ccccccccc"},
	}

	actual := parse(strings.NewReader(testCase))
	assert.Equal(t, expected, actual, "they should be equal")
}
