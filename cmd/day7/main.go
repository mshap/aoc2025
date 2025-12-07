package main

import (
	"aoc2025/internal/files"
	"fmt"
)

type point struct {
	start     bool
	splitter  bool
	processed bool
	beam      bool
}

func getStart(points [][]*point) (int, int) {
	for i, row := range points {
		for j, p := range row {
			if p.start {
				return i, j
			}
		}
	}
	return -1, -1
}

func countBeam(points [][]*point) int {
	total := 0
	for _, row := range points {
		for _, p := range row {
			if p.beam {
				total++
			}
		}
	}
	return total
}

func fireBeam(points [][]*point, row, col int) int {
	p := points[row][col]
	count := 0
	if p.processed {
		return 0
	}
	p.processed = true
	if !p.splitter {
		if p.beam || row >= len(points)-1 {
			return 0
		}
		return count + fireBeam(points, row+1, col)
	} else if p.splitter {
		// split := 0
		if col > 0 {
			count += fireBeam(points, row, col-1)
		}
		if col < len(points[row])-1 {
			count += fireBeam(points, row, col+1)
		}
		count++
	} else if row < len(points)-1 {
		return count + fireBeam(points, row+1, col)
	}

	return count
}

func main() {

	points := files.ProcessFile("../../internal/input/day7.txt", false, true, func(line string) ([]*point, error) {
		row := []*point{}

		for _, c := range line {
			p := point{}
			switch string(c) {
			case "S":
				p.start = true
			case "^":
				p.splitter = true
			}
			row = append(row, &p)
		}
		return row, nil
	})
	x, y := getStart(points)

	total := fireBeam(points, x+1, y)
	fmt.Printf("Total beams split: %d\n", total)
	fmt.Printf("Total beams counted: %d\n", countBeam(points))
}
