package main

import (
	"aoc2025/internal/files"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type list struct {
	start  int
	end    int
	remove bool
}

// func (l list) containsItem(id int) bool {
// 	return id >= l.start && id <= l.end
// }

type pantry struct {
	ranges []list
	// ingredients []int
}

func (p *pantry) rebuild() {
	slices.SortFunc(p.ranges, func(a, b list) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		}
		return 0
	})

	for i := 1; i < len(p.ranges); i++ {
		if p.ranges[i].start <= p.ranges[i-1].end {
			p.ranges[i].start = p.ranges[i-1].start
			if p.ranges[i-1].end > p.ranges[i].end {
				p.ranges[i].end = p.ranges[i-1].end
			}
			p.ranges[i-1].remove = true
		}
	}

	p.ranges = slices.DeleteFunc(p.ranges, func(r list) bool {
		return r.remove
	})
}

func (p pantry) getFreshItems() int {
	count := 0
	for _, r := range p.ranges {
		count += (r.end - r.start + 1)
	}

	// for _, id := range p.ingredients {
	// 	for _, r := range p.ranges {
	// 		if r.containsItem(id) && !slices.Contains(freshItems, id) {
	// 			freshItems = append(freshItems, id)
	// 			break
	// 		}
	// 	}
	// }
	return count
}

func main() {
	warehouse := pantry{}
	hitBlank := false
	files.ProcessFile("../../internal/input/day5.txt", true, true, func(line string) (pantry, error) {
		if line == "" {
			hitBlank = true
			return warehouse, nil
		}

		if !hitBlank {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			warehouse.ranges = append(warehouse.ranges, list{start: start, end: end})
		}
		// else {
		// 	val, _ := strconv.Atoi(line)
		// 	warehouse.ingredients = append(warehouse.ingredients, val)
		// }
		return warehouse, nil
	})
	warehouse.rebuild()
	// fmt.Printf("My warehouse %v\n", warehouse)

	fmt.Printf("%d Fresh items:\n", warehouse.getFreshItems())
}
