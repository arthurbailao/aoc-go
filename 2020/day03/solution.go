package day03

import (
	"bufio"
	"io"
	"strings"
)

// Solution solves dayX problems
type Solution struct {
	Matrix [][]string
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{Matrix: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	var trees int

	c := 3
	for l := 1; l < len(s.Matrix); l++ {
		item := s.Matrix[l][c]

		if item == "#" {
			trees++

		}

		c = (c + 3) % len(s.Matrix[l])

	}

	return trees, nil
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	trees := 1
	trees = trees * walkAndCount(s.Matrix, 1, 1)
	trees = trees * walkAndCount(s.Matrix, 3, 1)
	trees = trees * walkAndCount(s.Matrix, 5, 1)
	trees = trees * walkAndCount(s.Matrix, 7, 1)
	trees = trees * walkAndCount(s.Matrix, 1, 2)

	return trees, nil
}

func parse(input io.Reader) [][]string {
	var matrix [][]string

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, "")
		matrix = append(matrix, chars)

	}

	return matrix
}

func walkAndCount(matrix [][]string, right, down int) int {
	var trees int
	c := right
	for l := down; l < len(matrix); l = l + down {
		item := matrix[l][c]

		if item == "#" {
			trees++

		}

		c = (c + right) % len(matrix[l])

	}

	return trees
}
