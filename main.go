package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var puzzleInput string = fetchInput("input.txt")
	var parsedInput []string = strings.Split(puzzleInput, "\n")

	fmt.Printf("Total score: %d\n", calTotalScore(parsedInput))
}

func fetchInput(fn string) (input string) {
	content, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	input = string(content)
	input = strings.Trim(input, "\n")
	return
}

func calTotalScore(guide []string) int {

	var total int = 0
	for _, g := range guide {
		total += calSingleRoundScore(g)
	}
	return total
}

func calSingleRoundScore(roundGuide string) int {
	/*
		!Rules
		* A, X -> Rock (1)
		* B, Y -> Paper (2)
		* C, Z -> Scissors (3)
	*/
	var chosen []string = strings.Split(roundGuide, " ")
	var score int = 0

	if chosen[1] == "X" {
		score += 1
	} else if chosen[1] == "Y" {
		score += 2
	} else if chosen[1] == "Z" {
		score += 3
	}

	if (chosen[1] == "X" && chosen[0] == "C") ||
		(chosen[1] == "Y" && chosen[0] == "A") ||
		(chosen[1] == "Z" && chosen[0] == "B") {
		// win
		score += 6
	} else if (chosen[1] == "X" && chosen[0] == "A") ||
		(chosen[1] == "Y" && chosen[0] == "B") ||
		(chosen[1] == "Z" && chosen[0] == "C") {
		// draw
		score += 3
	} else {
		score += 0
	}

	return score
}
