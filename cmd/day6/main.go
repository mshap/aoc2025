package main

import (
	"aoc2025/internal/files"
	"fmt"
	"slices"
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
	nums []int
	op   operation
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

func main() {

	vals := files.ProcessFile("../../internal/input/day6.txt", func(line string) ([]string, error) {
		parts := strings.Split(line, " ")

		for i := 0; i < len(parts); i++ {
			parts[i] = strings.TrimSpace(parts[i])
		}

		return slices.DeleteFunc(parts, func(s string) bool {
			return s == ""
		}), nil
	})

	size := len(vals[0])
	maths := make([]mathProblem, size)

	for i := 0; i < len(vals); i++ {
		for j := 0; j < size; j++ {
			if i < len(vals)-1 {
				num, _ := strconv.Atoi(vals[i][j])
				maths[j].nums = append(maths[j].nums, num)
			} else {
				maths[j].op = operation(vals[i][j])
			}

		}
	}

	fmt.Printf("%v\n", maths)
	total := 0
	for _, v := range maths {
		total += v.calculate()
	}
	fmt.Printf("Total: %d\n", total)
}
