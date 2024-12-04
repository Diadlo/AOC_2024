package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
s  s  s
 a a s
  mmm
samxmas
  mmm
 a a a
s  s  s
*/

func addChar(data [][]byte, row int, col int, word *string) {
	if !(0 <= row && row < len(data) && 0 <= col && col < len(data[0])) {
		return
	}

	char := data[row][col]
	*word += string(char)
}

// Search for the word which is star at position (startRow, startCol) and goes in any direction
func countWords(data [][]byte, startRow int, startCol int, word string) int {
	// 4 directions for direct order. 4 other will be covered with reverse word
	var horizontal string
	var vertical string
	var mainDiag string
	var secondDiag string

	for i := range len(word) {
		addChar(data, startRow, startCol+i, &horizontal)
		addChar(data, startRow+i, startCol, &vertical)
		addChar(data, startRow+i, startCol+i, &mainDiag)
		addChar(data, startRow-i, startCol+i, &secondDiag)
	}

	count := 0
	if horizontal == word {
		count += 1
	}

	if vertical == word {
		count += 1
	}

	if mainDiag == word {
		count += 1
	}

	if secondDiag == word {
		count += 1
	}

	return count
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("Read")
	}

	lines := strings.Split(string(input), "\n")
	data := make([][]byte, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		data[i] = make([]byte, len(line))
		for j := range line {
			data[i][j] = line[j]
		}
	}

	count := 0
	for i := range data {
		for j := range data[i] {
			count += countWords(data, i, j, "XMAS")
			count += countWords(data, i, j, "SAMX")
		}
	}

	fmt.Printf("Count: %v\n", count)
}
