package day05

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strings"
)

type boardingPass struct {
	FrontBack, RightLeft []string
}

var fblrMap = map[string]int{
	"F": 1,
	"B": 0,
	"L": 1,
	"R": 0,
}

// Solution solves dayX problems
type Solution struct {
	BoardingPasses []boardingPass
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{BoardingPasses: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	var maxSeatID int

	for _, bp := range s.BoardingPasses {
		seatID := binarySpacePartitioning(bp.FrontBack, 128)*8 + binarySpacePartitioning(bp.RightLeft, 8)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID, nil
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	var seatIDs []int

	for _, bp := range s.BoardingPasses {
		seatID := binarySpacePartitioning(bp.FrontBack, 128)*8 + binarySpacePartitioning(bp.RightLeft, 8)
		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)

	for i := 0; i <= len(seatIDs); i++ {
		if (seatIDs[i+1] - seatIDs[i]) > 1 {
			return seatIDs[i] + 1, nil
		}
	}

	return -1, errors.New("seat not found")
}

func parse(input io.Reader) []boardingPass {
	var boardingPasses []boardingPass
	s := bufio.NewScanner(input)

	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, "")

		pass := boardingPass{}
		for i := 0; i < len(chars); i++ {
			if i < 7 {
				pass.FrontBack = append(pass.FrontBack, chars[i])
			} else {
				pass.RightLeft = append(pass.RightLeft, chars[i])
			}
		}

		boardingPasses = append(boardingPasses, pass)

	}

	return boardingPasses
}

func binarySpacePartitioning(input []string, size int) int {
	bounds := []int{0, size - 1}
	for _, c := range input {
		diff := (bounds[1] - bounds[0] + 1) / 2
		op := fblrMap[c]

		if op == 1 {
			diff = diff * -1
		}

		bounds[op] = bounds[op] + diff
	}

	if bounds[0] != bounds[1] {
		panic("failed to calculte bounds")
	}

	return bounds[0]
}
