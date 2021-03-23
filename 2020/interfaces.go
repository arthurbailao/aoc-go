package y2020

// Solver implementation solves a day problem
type Solver interface {
	SolveFirst() (int, error)
	SolveSecond() (int, error)
}
