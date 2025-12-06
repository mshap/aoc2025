package main

import (
	"aoc2025/internal/files"
	"fmt"
)

type space struct {
	contents int
	clear    bool
}

type warehouse struct {
	grid [][]space
}

func (w *warehouse) clear() bool {
	changed := false
	for i := range w.grid {
		for j := range w.grid[i] {
			if w.grid[i][j].clear {
				w.grid[i][j].contents = 0
				w.grid[i][j].clear = false
				changed = true
			}
		}
	}
	return changed
}

func (w *warehouse) getAdjacentCount(x, y int) int {
	if w.grid[x][y].contents == 0 {
		return 0
	}
	count := 0
	if y > 0 && w.grid[x][y-1].contents == 1 {
		count++
	}
	if y < len(w.grid[x])-1 && w.grid[x][y+1].contents == 1 {
		count++
	}

	return count
}

func (w *warehouse) getAbove(row, start, end int) int {
	if row < 0 {
		return 0
	} else if row >= len(w.grid) {
		return 0
	}

	if start < 0 {
		start = 0
	}
	if end > len(w.grid[row])-1 {
		end = len(w.grid[row]) - 1
	}
	count := 0
	for i := start; i <= end; i++ {
		if w.grid[row][i].contents == 1 {
			count++
		}
	}

	return count
}

// func (w *warehouse) getBelow(row, start, end int) int {
// 	if row < 0 {
// 		return 0
// 	} else if row > len(w.grid) {
// 		return 0
// 	}

// 	return 1
// }

func main() {
	rows := files.ProcessFile("../../internal/input/day4.txt", func(line string) ([]space, error) {
		var row []space
		for _, c := range line {
			var filled int
			if c == '@' {
				filled = 1
			}
			row = append(row, space{contents: filled})
		}
		return row, nil
	})
	warehouse := warehouse{grid: rows}
	rolls := 0
	repeat := true

	for repeat {
		for i, v := range warehouse.grid {
			for j := range v {
				if warehouse.grid[i][j].contents == 0 {
					continue
				}
				above := warehouse.getAbove(i-1, j-1, j+1)
				below := warehouse.getAbove(i+1, j-1, j+1)
				adj := warehouse.getAdjacentCount(i, j)
				total := above + below + adj
				if total < 4 {
					rolls++
					warehouse.grid[i][j].clear = true
				}

				// fmt.Printf("Grid %d-%d Above: %d, Below: %d, Adjacent: %d, Total: %d\n", i, j, above, below, adj, total)
			}
		}
		repeat = warehouse.clear()
	}

	fmt.Printf("Total rolls: %d\n", rolls)
	// fmt.Printf("%+v\n", warehouse)
}
