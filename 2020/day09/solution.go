package day09

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strconv"
)

// Solution solves day09 problems
type Solution struct {
	entries []int
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{entries: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	return s.SolveFirstWithParam(25)
}

// SolveFirstWithParam solves the first problem, receiving the size as first arg
func (s Solution) SolveFirstWithParam(size int) (int, error) {
	for i := size; i < len(s.entries); i++ {
		if !twoSum(s.entries[(i-size):i], s.entries[i]) {
			return s.entries[i], nil
		}
	}
	return -1, errors.New("solution not found")

}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	return s.SolveSecondWithParam(25)
}

// SolveSecondWithParam solves the first problem, receiving the size as first arg
func (s Solution) SolveSecondWithParam(size int) (int, error) {
	target, err := s.SolveFirstWithParam(size)
	if err != nil {
		return -1, err
	}

	var i, j, sum int
	for {

		if sum == target {
			break
		}

		if j >= len(s.entries) {
			return -1, errors.New("solution not found")
		}

		if sum < target {
			sum += s.entries[j]
			j++
		}

		if sum > target {
			sum -= s.entries[i]
			i++
		}
	}

	values := s.entries[i:j]
	sort.Ints(values)
	result := values[0] + values[len(values)-1]
	return result, nil

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

func twoSum(input []int, target int) bool {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == target {
				return true
			}
		}
	}

	return false
}
