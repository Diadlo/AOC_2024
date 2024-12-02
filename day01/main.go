package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("Failed to read")
	}

	lines := strings.Split(string(input), "\n")
	list1 := make([]int, len(lines)-1)
	list2 := make([]int, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		words := strings.Split(line, "   ")
		list1[i], err = strconv.Atoi(words[0])
		if err != nil {
			panicf("Failed to parse %v", words[0])
		}

		list2[i], err = strconv.Atoi(words[1])
		if err != nil {
			panicf("Failed to parse %v", words[1])
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	result := 0
	for i := range list1 {
		diff := abs(list1[i] - list2[i])
		result += diff
	}

	fmt.Printf("Part 1: %v\n", result)

	result = 0
	j := 0
	// Suboptimal solution because same number on the left side counted twice
	for i := range list1 {
		left := list1[i]
		for j < len(list2) && list2[j] < left {
			j += 1
		}

		save_j := j
		count := 0
		for j < len(list2) && list2[j] == left {
			j += 1
			count += 1
		}
		j = save_j

		result += left * count
	}

	fmt.Printf("Result2: %v\n", result)
}
