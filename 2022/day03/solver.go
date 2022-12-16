package day03

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func CharToPriority(char rune) int {
	if char >= 97 && char <= 122 {
		return int(char - 96)
	} else if char >= 65 && char <= 90 {
		return int(char - 38)
	} else {
		panic("invalid input: " + string(char))
	}
}

func Solve(rawInput []string) (string, string) {
	firstSolution := solvePart1(rawInput)
	secondSolution := solvePart2(rawInput)

	return fmt.Sprint(firstSolution), fmt.Sprint(secondSolution)
}

func solvePart2(input []string) (ouput int) {
	for i := 0; i < len(input); i += 3 {
		firstRucksack := input[i]
		secondRucksack := input[i+1]
		thirdRucksack := input[i+2]
		for _, char := range []rune(firstRucksack) {
			if strings.ContainsRune(secondRucksack, char) && strings.ContainsRune(thirdRucksack, char) {

				ouput += CharToPriority(char)
				break
			}
		}
	}
	return
}

func solvePart1(input []string) (output int) {
	var commonItems []rune
	for _, inputRow := range input {
		firstCompartment, secondCompartment := utils.HalveString(inputRow)
		for _, char := range []rune(firstCompartment) {
			if strings.ContainsRune(secondCompartment, char) {
				commonItems = append(commonItems, char)
				break
			}
		}
	}
	for _, item := range commonItems {
		output += CharToPriority(item)
	}
	return output
}
