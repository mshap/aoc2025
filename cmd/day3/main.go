package main

import (
	"aoc2025/internal/files"
	"fmt"
	"strconv"
)

type batteryBank struct {
	cells []int
}

func (b batteryBank) getMax(start, end int) (int, int) {
	max := 0
	index := 0
	for i := start; i < end; i++ {
		if b.cells[i] > max {
			max = b.cells[i]
			index = i
		}
	}
	return max, index
}

func processBank(bank batteryBank, cells []int, start, length int) []int {
	if len(cells) == length {
		return cells
	} else if start == len(bank.cells)-1 {
		cells = append(cells, bank.cells[start])
		return cells
	}

	end := len(bank.cells) - length + 1 + len(cells)
	// if start == 0 {
	// 	end -= (length - 1)
	// }
	val, index := bank.getMax(start, end)
	// second, _ := bank.getMax(index+1, len(bank.cells))

	return processBank(bank, append(cells, val), index+1, length)
}

func calcVoltage(cells []int) int {
	total := 0
	place := 1
	for i := len(cells) - 1; i >= 0; i-- {
		total += cells[i] * place
		place *= 10
	}
	return total
}

func main() {
	banks := files.ProcessFile("../../internal/input/day3.txt", false, true, func(line string) (batteryBank, error) {
		ints := make([]int, 0, len(line))
		for i := 0; i < len(line); i++ {
			volt, _ := strconv.Atoi(string(line[i]))
			ints = append(ints, volt)
		}
		return batteryBank{cells: ints}, nil
	})
	total := 0
	for _, v := range banks {
		res := processBank(v, []int{}, 0, 12)
		fmt.Printf("batteries: %v\n", res)
		total += calcVoltage(res)
	}
	fmt.Printf("Total Voltage: %d\n", total)
}
