package main

import (
	"aoc2022go/utils"
	"errors"
	"fmt"
	"strings"
)

func main() {
	puzzleInput, _ := utils.FetchInput("input.txt")

	rucksacks := strings.Split(puzzleInput, "\n")
	errors := findErrorInRucksacks(rucksacks)
	sum := sumOfPriorities(errors)
	fmt.Println(sum)
}

func findErrorInRucksacks(rucksacks []string) []byte {
	var errors []byte
	for _, r := range rucksacks {
		detected := detectItemTypeError(r)
		if detected != 0 {
			errors = append(errors, detected)
		}
	}
	return errors
}

func sumOfPriorities(itemTypes []byte) int {
	sum := 0
	for _, i := range itemTypes {
		priority, _ := convItemTypeToPriority(i)
		sum += priority
	}
	return sum
}

func detectItemTypeError(rucksack string) byte {
	// use simple intersection algorithm with O(n^2)
	var items []byte = []byte(rucksack)
	first := items[:(len(items) / 2)]
	second := items[(len(items) / 2):]
	// fmt.Printf("first: %d, second: %d\n", len(first), len(second))
	for _, f := range first {
		for _, s := range second {
			if f == s {
				return f
			}
		}
	}
	return 0
}

func convItemTypeToPriority(itemType byte) (int, error) {
	const (
		StartUppercase = 65
		StopUppercase  = 90
		StartLowercase = 97
		StopLowercase  = 122
	)
	var priority int
	if i := int(itemType); i >= StartUppercase && i <= StopUppercase {
		priority = i - (StartUppercase - 1) + 26 // add 26 for more priority
	} else if i >= StartLowercase && i <= StopLowercase {
		priority = i - (StartLowercase - 1)
	} else {
		return 0, errors.New("invalid item type format, item type must be a-z or A-Z")
	}
	return priority, nil
}
