package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	first  int
	second int
}

func parseInput(input string) ([]Rule, [][]int) {
	lines := strings.Split(string(input), "\n")
	rules := make([]Rule, 0)
	updates := make([][]int, 0)
	workOnRules := true
	for _, line := range lines {
		if line == "" {
			workOnRules = false
			continue
		}

		if workOnRules {
			values := strings.Split(line, "|")
			left, err := strconv.Atoi(values[0])
			if err != nil {
				panic("Parse")
			}
			right, err := strconv.Atoi(values[1])
			if err != nil {
				panic("Parse")
			}
			rules = append(rules, Rule{left, right})
		} else {
			strValues := strings.Split(line, ",")
			values := make([]int, len(strValues))
			for i, strValue := range strValues {
				intValue, err := strconv.Atoi(strValue)
				if err != nil {
					panic("Parse")
				}
				values[i] = intValue
			}
			updates = append(updates, values)
		}
	}

	return rules, updates
}

func updateIsCorrect(rulesForPages [][]Rule, update []int) bool {
	setOfPages := make(map[int]bool)
	for _, page := range update {
		setOfPages[page] = true
	}

	for i, page := range update {
		rightSet := make(map[int]bool)
		for _, p := range update[i+1:] {
			rightSet[p] = true
		}

		rulesForPage := rulesForPages[page]
		for _, rule := range rulesForPage {
			if !setOfPages[rule.first] || !setOfPages[rule.second] {
				// One of the page from rule is missing - skip
				continue
			}

			if page == rule.first {
				if !rightSet[rule.second] {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("Read")
	}

	rules, updates := parseInput(string(input))

	rulesForPages := make([][]Rule, 101)
	for i := range 100 {
		rulesForPages[i] = make([]Rule, 0)
	}

	for _, rule := range rules {
		rulesForPages[rule.first] = append(rulesForPages[rule.first], rule)
		rulesForPages[rule.second] = append(rulesForPages[rule.second], rule)
	}

	correctUpdates := make([][]int, 0)
	incorrectUpdates := make([][]int, 0)
	for _, update := range updates {
		if updateIsCorrect(rulesForPages, update) {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	//fmt.Printf("Correct: %v\n", correctUpdates)
	fmt.Printf("Incorrect: %v\n", incorrectUpdates)

	result := 0
	for _, update := range correctUpdates {
		result += update[len(update)/2]
	}

	fmt.Printf("result: %v", result)
}
