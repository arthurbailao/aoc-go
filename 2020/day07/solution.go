package day07

import (
	"bufio"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	Bag      string
	Contents map[string]int
}

func (r1 rule) equal(r2 rule) bool {
	if r1.Bag != r2.Bag {
		return false

	}

	return reflect.DeepEqual(r1.Contents, r2.Contents)

}

// Solution solves day01 problems
type Solution struct {
	rules []rule
}

// New builds the Solution for the day 7
func New(input io.Reader) Solution {
	return Solution{rules: parse(input)}
}

// SolveFirst solves the first problem
func (s Solution) SolveFirst() (int, error) {

	graph := map[string][]string{}

	for _, rule := range s.rules {
		for k := range rule.Contents {
			graph[k] = append(graph[k], rule.Bag)
		}
	}

	result := findBags(graph, "shiny gold")
	return len(result), nil
}

// SolveSecond solves the second problem
func (s Solution) SolveSecond() (int, error) {
	graph := map[string]map[string]int{}

	for _, rule := range s.rules {
		graph[rule.Bag] = rule.Contents
	}

	return countBags(graph, "shiny gold"), nil
}

func parse(input io.Reader) []rule {
	var rules []rule
	s := bufio.NewScanner(input)

	re := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)

	for s.Scan() {
		parts := strings.Split(s.Text(), "bags contain")
		bag := strings.TrimSpace(parts[0])

		if parts[1] == " no other bags." {
			rules = append(rules, rule{bag, map[string]int{}})
			continue

		}

		matches := re.FindAllStringSubmatch(parts[1], -1)
		contents := map[string]int{}

		for _, m := range matches {
			count, _ := strconv.Atoi(m[1])
			contents[m[2]] = count

		}

		rules = append(rules, rule{bag, contents})

	}
	return rules

}

func findBags(graph map[string][]string, bag string) map[string]bool {
	bags := map[string]bool{}

	for _, c := range graph[bag] {
		bags[c] = true
		for k := range findBags(graph, c) {
			bags[k] = true

		}

	}
	return bags

}

func countBags(graph map[string]map[string]int, bag string) int {
	var total int

	for k, v := range graph[bag] {
		total = total + v + v*countBags(graph, k)

	}

	return total

}
