package main

import (
	"aoc2025/internal/files"
	"fmt"
	"strconv"
	"strings"
)

type machine struct {
	goal       []bool
	indicators []bool
	buttons    [][]int
	pushes     int
	best       int
}

func (m machine) getPushes() int {
	m.pushButton(m.buttons[0])
	return m.best
}

func (m *machine) reset() {
	m.indicators = make([]bool, len(m.goal))
	m.pushes = 0
}

func (m *machine) pushButton(button []int) {
	for _, v := range button {
		m.indicators[v] = true
	}
	m.pushes++
}

func main() {
	machines := files.ProcessFile("../../internal/input/day10.txt", false, true, func(line string) (machine, error) {
		parts := strings.Split(line, " ")

		goal := parseGoal(parts[0][1 : len(parts[0])-1])
		buttons := [][]int{}
		for i := 1; i < len(parts); i++ {
			buttons = append(buttons, parseButton(parts[i]))
		}
		m := machine{goal: goal, buttons: buttons}
		m.reset()

		return m, nil
	})

	fmt.Printf("%+v\n", machines)

	total := 0
	for _, m := range machines {
		total += m.getPushes()
	}
	fmt.Printf("Total pushes: %d\n", total)
}

func parseGoal(s string) []bool {
	goal := []bool{}
	for i := range len(s) {
		if string(s[i]) == "#" {
			goal = append(goal, true)
		} else {
			goal = append(goal, false)
		}
	}
	return goal
}

func parseButton(s string) []int {
	parts := strings.Split(s[1:len(s)-1], ",")
	button := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		button[i], _ = strconv.Atoi(parts[i])
	}
	return button
}
