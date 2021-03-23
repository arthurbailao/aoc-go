/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cmd defines all Cobra commands.
package cmd

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	y2020 "github.com/arthurbailao/aoc/2020"
	"github.com/arthurbailao/aoc/2020/day01"
	"github.com/arthurbailao/aoc/2020/day02"
	"github.com/arthurbailao/aoc/2020/day03"
	"github.com/arthurbailao/aoc/2020/day04"
	"github.com/arthurbailao/aoc/2020/day05"
	"github.com/arthurbailao/aoc/2020/day06"
	"github.com/arthurbailao/aoc/2020/day07"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// y2020Cmd represents the 2020 command
var y2020Cmd = &cobra.Command{
	Use:   "2020",
	Short: "Executes 2020 solves",
	Long: `2020 is for solving any problem from this year. Just choose the day 
you want and provide an input file at the first argument.

For example, 'aoc 2020 day07 input.txt' will solve 
the 7th day problem of the year 2020.`,
}

var solutions = map[string]func(io.Reader) y2020.Solver{
	"day01": func(input io.Reader) y2020.Solver { return day01.New(input) },
	"day02": func(input io.Reader) y2020.Solver { return day02.New(input) },
	"day03": func(input io.Reader) y2020.Solver { return day03.New(input) },
	"day04": func(input io.Reader) y2020.Solver { return day04.New(input) },
	"day05": func(input io.Reader) y2020.Solver { return day05.New(input) },
	"day06": func(input io.Reader) y2020.Solver { return day06.New(input) },
	"day07": func(input io.Reader) y2020.Solver { return day07.New(input) },
}

func init() {
	rootCmd.AddCommand(y2020Cmd)
	re := regexp.MustCompile("[0-9]+")

	for cmdName := range solutions {
		match := re.FindString(cmdName)
		day, _ := strconv.Atoi(match)

		subcommand := cobra.Command{
			Use:   fmt.Sprintf("%s inputfile", cmdName),
			Short: fmt.Sprintf("solves https://adventofcode.com/2020/day/%d", day),
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				file, err := os.Open(args[0])
				if err != nil {
					return errors.Wrap(err, "open file")
				}

				solver := solutions[cmd.Name()](file)

				r, err := solver.SolveFirst()
				if err != nil {
					return errors.Wrap(err, "failed to solve the first part")
				}

				fmt.Printf("part one: %d\n", r)

				r, err = solver.SolveSecond()
				if err != nil {
					return errors.Wrap(err, "failed to solve the second part")
				}

				fmt.Printf("part two: %d\n", r)

				return nil
			},
		}
		y2020Cmd.AddCommand(&subcommand)
	}
}
