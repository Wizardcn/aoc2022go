package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"log"
)

func main() {
	// read puzzle input from text file
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var puzzleInput string = string(content)
	// fmt.Println(puzzleInput)
	fmt.Printf("total calories that Elf carrying: %d", findMaxCal(puzzleInput))
}

func findMaxCal(elfCals string) int {

	var sumCals []int
	var calElves []string = strings.Split(elfCals, "\n\n")
	fmt.Printf("number of elves: %d\n", len(calElves))

	for _, cal := range calElves {
		calCount := 0
		cal := strings.Trim(cal, "\n")
		calList := strings.Split(cal, "\n")
		for _, c := range calList {
			cInt, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println(calList, len(calList))
				fmt.Printf("Cannot cast %s to int type", c)
				return -1
			}
			calCount += cInt
		}
		sumCals = append(sumCals, calCount)
	}
	return maxIntSlice(sumCals)
}

func maxIntSlice(sl []int) int {
	var max int = sl[0]
	for _, i := range sl {
		if i > max {
			max = i
		}
	}
	return max
}
