package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("Read")
	}

	r := regexp.MustCompile(`(mul\(([0-9][0-9]?[0-9]?),([0-9][0-9]?[0-9]?)\)|do\(\)|don't\(\))`)
	matches := r.FindAllStringSubmatch(string(input), -1)

	enabled := true
	result := 0
	for _, match := range matches {
		switch {
		case strings.HasPrefix(match[0], "mul(") && enabled:
			match := match[1:]
			left, err := strconv.Atoi(match[1])
			if err != nil {
				panic("atoi")
			}

			right, err := strconv.Atoi(match[2])
			if err != nil {
				panic("atoi")
			}

			result += left * right
		case match[0] == "do()":
			enabled = true
		case match[0] == "don't()":
			enabled = false
		}
	}

	fmt.Printf("%v\n", result)
}
