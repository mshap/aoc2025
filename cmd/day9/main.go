package main

import (
	"aoc2025/internal/files"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type color string

const (
	RED   color = "red"
	GREEN color = "green"
)

type tile struct {
	col, row int
	color    color
}

func main() {

	reds := files.ProcessFile("../../internal/input/day9.txt", false, true, func(line string) (tile, error) {
		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[0])
		return tile{row: x, col: y, color: RED}, nil
	})

	// slices.SortFunc(reds, func(a, b tile) int {
	// 	if a.row < b.row {
	// 		return -1
	// 	} else if a.row > b.row {
	// 		return 1
	// 	} else {
	// 		if a.col < b.col {
	// 			return -1
	// 		} else if a.col > b.col {
	// 			return 1
	// 		} else {
	// 			return 0
	// 		}
	// 	}
	// })
	// fmt.Printf("%+v\n", reds)
	max := 0.0

	for idx, tile := range reds {
		last := tile
		for i := 1; i < len(reds); i++ {
			if i == idx {
				continue
			}
			cur := reds[i]
			colDist := cur.col - last.col + 1
			rowDist := cur.row - last.row + 1
			size := math.Abs(float64(colDist * rowDist))
			if size > max {
				max = size
				// last = cur
			}
		}
	}

	fmt.Printf("Max size: %.0f\n", max)

}
