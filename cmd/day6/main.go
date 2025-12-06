package main

import (
	"aoc2025/internal/files"
	"fmt"
	"strconv"
	"strings"
)

type operation string

const (
	add      operation = "+"
	subtract operation = "-"
	multiply operation = "*"
	divide   operation = "/"
)

type mathProblem struct {
	numsStr []string
	nums    []int
	op      operation
	length  int
}

func (m mathProblem) calculate() int {
	total := m.nums[0]

	for i := 1; i < len(m.nums); i++ {
		switch m.op {
		case add:
			total += m.nums[i]
		case subtract:
			total -= m.nums[i]
		case multiply:
			total *= m.nums[i]
		case divide:
			total /= m.nums[i]
		}
	}
	return total
}

func (m *mathProblem) convert() {
	for i := 0; i < m.length; i++ {
		var str string
		for j, s := range m.numsStr {
			if j == len(m.numsStr)-1 {
				continue
			}
			str += string(s[i])
		}
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		m.nums = append(m.nums, num)
	}
}

func createOperations(opLine string) []*mathProblem {
	parts := strings.Split(opLine, " ")
	var problems []*mathProblem

	var op operation
	var problem *mathProblem
	for i := 0; i < len(parts); i++ {
		if parts[i] != "" {
			op = operation(parts[i])
			problem = &mathProblem{op: op}
			problems = append(problems, problem)
		}
		problem.length = problem.length + 1
	}
	return problems
}

func main() {

	lines := files.ProcessFile("../../internal/input/day6.txt", false, false, func(line string) (string, error) {
		return line, nil
	})

	opLine := lines[len(lines)-1]

	maths := createOperations(opLine)
	for _, line := range lines {
		x := 0
		for _, p := range maths {
			p.numsStr = append(p.numsStr, line[x:x+p.length])
			x += (p.length + 1)
		}
	}

	total := 0
	for _, v := range maths {
		v.convert()
		total += v.calculate()
	}
	fmt.Printf("Total: %d\n", total)
}
