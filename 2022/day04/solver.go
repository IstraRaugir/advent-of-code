package day04

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func Solve(rawInput []string) (string, string) {
	fullyContainedSectionPairs := 0
	overlappingSectionPairs := 0
	for _, inputRow := range rawInput {
		firstSection, secondSection := parseSectionPairs(inputRow)
		if sectionFullyContains(firstSection, secondSection) {
			fullyContainedSectionPairs++
		}
		if sectionsOverlap(firstSection, secondSection) {
			overlappingSectionPairs++
		}
	}
	return fmt.Sprint(fullyContainedSectionPairs), fmt.Sprint(overlappingSectionPairs)
}

func sectionsOverlap(first [2]int, second [2]int) bool {
	if first[0] > first[1] || second[0] > second[1] {
		panic(fmt.Sprint("Input was not ascending", first, second))
	}
	return first[1] >= second[0] && first[0] <= second[1]
}

func parseSectionPairs(inputRow string) ([2]int, [2]int) {
	var assignmentSet []int
	for _, rawAssignment := range strings.Split(inputRow, ",") {
		for _, rawSectionId := range strings.Split(rawAssignment, "-") {
			assignmentSet = append(assignmentSet, utils.ParseInt(rawSectionId))
		}
	}
	firstSection := [2]int{assignmentSet[0], assignmentSet[1]}
	secondSection := [2]int{assignmentSet[2], assignmentSet[3]}
	return firstSection, secondSection
}

func sectionFullyContains(first [2]int, second [2]int) bool {
	if first[0] > first[1] || second[0] > second[1] {
		panic(fmt.Sprint("Input was not ascending", first, second))
	}
	firstContainedInSecond := first[0] >= second[0] && first[1] <= second[1]
	secondContainedInFirst := second[0] >= first[0] && second[1] <= first[1]
	return firstContainedInSecond || secondContainedInFirst
}
