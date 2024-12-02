package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func panicf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	panic(s)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func reportIsSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	// The levels are either all increasing or all decreasing
	if report[0] == report[1] {
		return false
	}

	mustBeAscending := report[0] < report[1]

	for i := range len(report) - 1 {
		ascending := report[i] < report[i+1]
		descending := report[i] > report[i+1]
		diff := abs(report[i] - report[i+1])
		if (1 <= diff && diff <= 3) && (ascending && mustBeAscending || descending && !mustBeAscending) {
			continue
		}

		return false
	}

	return true
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panicf("Failed to read input")
	}

	lines := strings.Split(string(input), "\n")
	reports := make([][]int, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}

		elements := strings.Split(line, " ")
		reports[i] = make([]int, len(elements))
		for j, elem := range elements {
			reports[i][j], err = strconv.Atoi(elem)
			if err != nil {
				panicf("Failed to parse '%v' on line %v", elem, i)
			}
		}
	}

	count := 0
	for _, report := range reports {
		if reportIsSafe(report) {
			count += 1
		}
	}

	fmt.Printf("Count: %v", count)
}
