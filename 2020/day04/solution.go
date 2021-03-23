package day04

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

// Solution solves dayX problems
type Solution struct {
	Passports []passport
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{Passports: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	var valid int
	for _, p := range s.Passports {
		if len(p) >= 8 {
			valid++
		} else if len(p) >= 7 && p["cid"] == "" {
			valid++
		}
	}
	return valid, nil
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	validators := map[string]func(string) bool{
		"byr": intervalValidator(1920, 2002),
		"iyr": intervalValidator(2010, 2020),
		"eyr": intervalValidator(2020, 2030),
		"hgt": heightValidator,
		"hcl": regexpValidator("#[0-9a-f]{6}"),
		"ecl": regexpValidator("(amb|blu|brn|gry|grn|hzl|oth)"),
		"pid": regexpValidator("^[0-9]{9}$"),
		"cid": func(string) bool { return false },
	}

	var total int
	for _, p := range s.Passports {
		var valids []bool
		for k, v := range p {
			valids = append(valids, validators[k](v))
		}

		var trues int
		for _, b := range valids {
			if b {
				trues++
			}
		}

		if trues == 7 {
			total++
		}

	}

	return total, nil
}

func parse(input io.Reader) []passport {
	var passports []passport

	current := passport{}
	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		words := strings.Split(line, " ")

		if len(line) == 0 {
			passports = append(passports, current)
			current = make(passport)
			continue
		}

		for _, w := range words {
			kv := strings.Split(w, ":")
			current[kv[0]] = kv[1]
		}
	}

	if len(current) > 0 {
		passports = append(passports, current)

	}

	return passports
}

func intervalValidator(min, max int) func(string) bool {
	return func(s string) bool {

		i, err := strconv.Atoi(s)
		if err != nil {
			return false

		}
		return i >= min && i <= max
	}
}

func regexpValidator(r string) func(string) bool {
	return func(s string) bool {
		re := regexp.MustCompile(r)
		return re.MatchString(s)
	}
}

func heightValidator(s string) bool {
	switch symbol := s[len(s)-2 : len(s)]; symbol {
	case "cm":
		return intervalValidator(150, 193)(s[:len(s)-2])
	case "in":
		return intervalValidator(59, 76)(s[:len(s)-2])
	default:
		return false
	}
}
