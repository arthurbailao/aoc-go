package day08

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/arthurbailao/aoc/2020/stack"
)

type instruction struct {
	Cmd string
	Val int
}

type program struct {
	Instructions []instruction
	Acc          int
	Cmds         []int
	Stack        *stack.Stack
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
	prog := newProgram(s.Instructions)

	err := prog.Run()

	if err == nil {
		return -1, errors.New("no ininite loop found")
	}

	if err.Error() == "ininite loop" {
		return prog.Acc, nil
	}

	return -1, err
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	fixMap := map[string]string{
		"jmp": "nop",
		"nop": "jmp",
		"acc": "acc",
	}

	var testingInstructions []int
	for i, inst := range s.Instructions {
		if inst.Cmd != "acc" {
			testingInstructions = append(testingInstructions, i)
		}
	}

	for _, i := range testingInstructions {
		modified := make([]instruction, len(s.Instructions))
		copy(modified, s.Instructions)

		modified[i].Cmd = fixMap[modified[i].Cmd]

		prog := newProgram(modified)
		err := prog.Run()

		if err == nil {
			return prog.Acc, nil
		}
	}

	return -1, errors.New("could not fix the program")

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

func newProgram(instructions []instruction) program {
	prog := program{
		Instructions: instructions,
		Acc:          0,
		Cmds:         make([]int, len(instructions)),
		Stack:        stack.New(),
	}

	prog.Stack.Push(0)
	return prog
}

func (p *program) Run() error {
	for {
		pos := p.Stack.Peek().(int)
		if p.Cmds[pos] == 1 {
			return errors.New("ininite loop")
		}
		p.Cmds[pos]++

		inc := 1

		switch current := p.Instructions[pos]; current.Cmd {
		case "nop":
		case "acc":
			p.Acc = p.Acc + current.Val
		case "jmp":
			inc = current.Val
		default:
			return fmt.Errorf("cmd `%s` does not exists", current.Cmd)
		}

		if pos+inc == len(p.Instructions) {
			break
		}

		p.Stack.Push(pos + inc)
	}
	return nil
}
