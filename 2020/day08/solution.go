package day08

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type instruction struct {
	Cmd string
	Val int
}

type program struct {
	Instructions []instruction
	Acc          int
	Cmds         []int
}

// Solution solves day01 problems
type Solution struct {
	Instructions []instruction
}

// New builds the Solution for the day 1
func New(input io.Reader) Solution {
	return Solution{Instructions: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {
	prog := program{
		Instructions: s.Instructions,
		Acc:          0,
		Cmds:         make([]int, len(s.Instructions)),
	}

	err := prog.Run()

	if err == nil {
		return -1, errors.New("no ininite loop found")
	}

	if err.Error() == "ininite loop" {
		return prog.Acc, nil
	}

	return -1, errors.New("unknown error")
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	return -1, errors.New("not implemented")
}

func parse(input io.Reader) []instruction {
	var program []instruction

	s := bufio.NewScanner(input)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, " ")

		val, _ := strconv.Atoi(parts[1])

		program = append(program, instruction{Cmd: parts[0], Val: val})
	}

	return program
}

func (p *program) Run() error {
	var i int
	for {
		if p.Cmds[i] == 1 {
			return errors.New("ininite loop")
		}
		p.Cmds[i]++

		switch current := p.Instructions[i]; current.Cmd {
		case "nop":
			i++
		case "acc":
			p.Acc = p.Acc + current.Val
			i++
		case "jmp":
			i = i + current.Val
		default:
			return errors.New("cmd does not exists")
		}
	}
	return nil
}
