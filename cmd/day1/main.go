package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction string

type turn struct {
	dir    direction
	amount int
}

const (
	start int       = 50
	max   int       = 100
	left  direction = "L"
	right direction = "R"
)

var (
	zero    int
	current int = start
)

func getCommands() []turn {
	var turns []turn = []turn{}

	inputLines, err := readInputFile("../internal/input/day1.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return turns
	}

	fmt.Printf("Read %d lines from input file\n", len(inputLines))
	for _, line := range inputLines {
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error parsing amount: %v\n", err)
			return nil
		}

		command := turn{
			dir:    direction(line[0:1]),
			amount: amount,
		}
		turns = append(turns, command)
	}
	return turns
}

func processTurn(turn turn, count int) (int, int) {
	var dial int
	var multiplier int = turn.amount / max
	var turns = turn.amount % max
	switch turn.dir {
	case left:
		dial = count - turns
	case right:
		dial = count + turns
	}
	extra := multiplier

	if dial == 0 {
		return 0, 1 + extra
	} else if dial < 0 {
		dial = max + dial
		if count != 0 {
			extra++
		}
		// return max + count, 1
	} else if dial >= max {
		dial -= max
		// return count - max, 1
		extra++
	}

	return dial, extra
}

func main() {
	turns := getCommands()
	fmt.Printf("Current position: %d %d\n", current, 0)
	for _, turn := range turns {
		fmt.Printf("Turn: %v\n", turn)
		var extra int
		current, extra = processTurn(turn, current)

		if current == 0 && extra == 0 {
			zero++
		}

		zero += extra

		fmt.Printf("Current position: %d %d\n", current, extra)
	}

	fmt.Printf("Zero turns: %d\n", zero)
}

// readInputFile reads a text file and returns its lines as a slice of strings
func readInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" { // Skip empty lines
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
