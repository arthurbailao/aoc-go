package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type policy struct {
	Range  []int
	Letter string
}

type policyAndPassword struct {
	Policy   policy
	Password string
}

// Solution solves dayX problems
type Solution struct {
	Policies []policyAndPassword
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{Policies: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	var count int
	for _, input := range s.Policies {
		var charCount int
		for _, char := range input.Password {
			if string(char) == input.Policy.Letter {
				charCount++

			}

		}

		if charCount >= input.Policy.Range[0] && charCount <= input.Policy.Range[1] {
			count++
		}

	}

	return count, nil
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	var count int
	for _, input := range s.Policies {
		var sum int
		chars := strings.Split(input.Password, "")

		for _, n := range input.Policy.Range {
			if chars[n-1] == input.Policy.Letter {
				sum++

			}

		}

		if sum == 1 {
			count++

		}

	}

	return count, nil
}

func parse(input io.Reader) []policyAndPassword {
	var inputs []policyAndPassword

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		substrings := strings.Split(line, " ")

		minMax := strings.Split(substrings[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		letter := strings.Split(substrings[1], "")[0]
		p := policy{Range: []int{min, max}, Letter: letter}

		inputs = append(inputs, policyAndPassword{Policy: p, Password: substrings[2]})
	}

	return inputs
}
