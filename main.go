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
		* A -> Rock (1)
		* B -> Paper (2)
		* C -> Scissors (3)

		* X -> lose
		* Y -> draw
		* Z -> win


	*/
	var chosen []string = strings.Split(roundGuide, " ")
	var score int = 0

	if chosen[1] == "X" {
		score += 0
	} else if chosen[1] == "Y" {
		score += 3
	} else if chosen[1] == "Z" {
		score += 6
	}

	if (chosen[1] == "Z" && chosen[0] == "C") ||
		(chosen[1] == "X" && chosen[0] == "B") ||
		(chosen[1] == "Y" && chosen[0] == "A") {
		// rock
		score += 1
	} else if (chosen[1] == "X" && chosen[0] == "C") ||
		(chosen[1] == "Y" && chosen[0] == "B") ||
		(chosen[1] == "Z" && chosen[0] == "A") {
		// paper
		score += 2
	} else {
		// scissors
		score += 3
	}

	return score
}
