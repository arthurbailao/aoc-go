package day01

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

// Solution solves day01 problems
type Solution struct {
	entries []int
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{entries: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	for i := 0; i < len(s.entries)-1; i++ {
		for j := i + 1; j < len(s.entries); j++ {
			if s.entries[i]+s.entries[j] == 2020 {
				return s.entries[i] * s.entries[j], nil
			}
		}
	}

	return -1, errors.New("could not find two entries that sums 2020")
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	for i := 0; i < len(s.entries)-2; i++ {
		for j := i + 1; j < len(s.entries)-1; j++ {
			for k := j + 1; k < len(s.entries); k++ {
				if s.entries[i]+s.entries[j]+s.entries[k] == 2020 {
					return s.entries[i] * s.entries[j] * s.entries[k], nil

				}

			}

		}

	}
	return -1, errors.New("could not find three entries that sums 2020")
}

func parse(input io.Reader) []int {
	var numbers []int

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	return numbers
}
