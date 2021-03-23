package day04

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929
`

func TestParse(t *testing.T) {
	expected := []passport{
		passport{
			"ecl": "gry",
			"pid": "860033327",
			"eyr": "2020",
			"hcl": "#fffffd",
			"byr": "1937",
			"iyr": "2017",
			"cid": "147",
			"hgt": "183cm",
		},
		passport{
			"iyr": "2013",
			"ecl": "amb",
			"cid": "350",
			"eyr": "2023",
			"pid": "028048884",
			"hcl": "#cfa07d",
			"byr": "1929",
		},
	}

	actual := parse(strings.NewReader(input))
	assert.Equal(t, expected, actual, "they should be equal")
}
