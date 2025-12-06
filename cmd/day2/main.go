package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type productRange struct {
	start int
	end   int
}

func (p productRange) getDupeNbrs() []int {
	start := p.start
	end := p.end

	var dupeNumbers []int

	for i := start; i <= end; i++ {
		strI := strconv.Itoa(i)

		for group := 2; group <= len(strI); group++ {
			if isDupByLength(strI, group) {
				dupeNumbers = append(dupeNumbers, i)
				break
			}
		}

	}
	return dupeNumbers
}

func isDupByLength(s string, groups int) bool {
	if len(s)%groups != 0 {
		return false
	}
	var parts []int = make([]int, groups)

	size := len(s) / groups

	for i := 0; i < groups; i++ {
		start := i * size
		end := start + size
		if i == groups-1 {
			end = len(s)
		}
		part, _ := strconv.Atoi(s[start:end])
		parts[i] = part
	}

	match := parts[0]
	for i := 1; i < len(parts); i++ {
		if parts[i] != match {
			return false
		}
	}
	return true
}

func getProducts() []productRange {
	var products []productRange

	// Read the input file
	inputLines, err := readInputFile("../../internal/input/day2.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return products
	}

	fmt.Printf("Read %d lines from input file\n", len(inputLines))

	// Print first few lines as example
	for _, line := range inputLines {
		parts := strings.Split(line, ",")

		for _, v := range parts {
			ranges := strings.Split(v, "-")
			if len(ranges) == 2 {
				start, err := strconv.Atoi(ranges[0])
				if err != nil {
					continue
				}
				end, err := strconv.Atoi(ranges[1])
				if err != nil {
					continue
				}
				products = append(products, productRange{start: start, end: end})
			}
		}
	}

	return products
}

func main() {
	var dupes []int
	var total int
	products := getProducts()
	for _, p := range products {
		dupes = append(dupes, p.getDupeNbrs()...)
		fmt.Printf("Product range: %d - %d has %v duplicates\n", p.start, p.end, p.getDupeNbrs())

	}

	for _, d := range dupes {
		total += d
	}
	fmt.Printf("Total duplicates found: %d\n", total)
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
