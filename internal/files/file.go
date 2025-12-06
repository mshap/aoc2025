package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessFile[T any](fileName string, skip, trim bool, parser func(string) (T, error)) []T {

	inputLines, err := readInputFile(fileName, skip, trim)

	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return nil
	}

	fmt.Printf("Read %d lines from input file\n", len(inputLines))
	var ts []T = []T{}
	// Print first few lines as example
	for _, line := range inputLines {
		t, err := parser(line)
		if err == nil {
			ts = append(ts, t)
		}
	}

	return ts
}

func readInputFile(filename string, skipEmpty, trim bool) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if trim {
			line = strings.TrimSpace(line)
		}
		if !skipEmpty || line != "" { // Skip empty lines
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
