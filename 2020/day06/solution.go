package day06

import (
	"bufio"
	"io"
	"strings"
)

type group struct {
	Answers         map[string]int
	NumOfPassengers int
}

// Solution solves dayX problems
type Solution struct {
	Groups []group
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{Groups: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	var total int
	for _, g := range s.Groups {
		total = total + len(g.Answers)
	}

	return total, nil
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	var total int
	for _, g := range s.Groups {
		for _, v := range g.Answers {
			if v >= g.NumOfPassengers {
				total++
			}
		}
	}

	return total, nil
}

func parse(input io.Reader) []group {
	var groups []group
	var grp = group{Answers: make(map[string]int), NumOfPassengers: 0}

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()

		if line == "" {
			groups = append(groups, grp)

			grp = group{Answers: make(map[string]int), NumOfPassengers: 0}
			continue

		}

		grp.NumOfPassengers++
		answers := strings.Split(line, "")
		for _, a := range answers {
			grp.Answers[a]++
		}
	}

	if len(grp.Answers) > 0 {
		groups = append(groups, grp)

	}

	return groups
}
