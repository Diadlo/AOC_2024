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

func pairIsGood(a int, b int, mustBeAscending bool) bool {
	ascending := a < b
	descending := a > b
	diff := abs(a - b)

	return (1 <= diff && diff <= 3) &&
		(ascending && mustBeAscending || descending && !mustBeAscending)
}

func countErrors(report []int, mustBeAscending bool) int {
	count := 0
	for i := range len(report) - 1 {
		if pairIsGood(report[i], report[i+1], mustBeAscending) {
			// Great case
			continue
		}

		// One error is allowed, let's check if the same condition work if one element is skipped

		// If this is the last but one element, we can drop the next element (last) without any check
		if i == len(report)-2 {
			count += 1
			continue
		}

		if pairIsGood(report[i], report[i+2], mustBeAscending) {
			count += 1
			continue
		}

		// If it doesn't work even with the skipped element - report is bad.
		// Say that we have a lot of errors
		return 99
	}

	// If we faced only one bad element, it's ok
	return count
}

func CheckReport(report []int) bool {
	if len(report) < 2 {
		return true
	}

	return countErrors(report, true) <= 1 || countErrors(report, false) <= 1 ||
		countErrors(report[1:], true) == 0 || countErrors(report[1:], false) == 0
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
		fmt.Printf("%v -> %v\n", report, CheckReport(report))
		if CheckReport(report) {
			count += 1
		}
	}

	fmt.Printf("Count: %v", count)
}

// 424 - 500
