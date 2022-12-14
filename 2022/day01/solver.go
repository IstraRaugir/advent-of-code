package day01

import (
	"aoc2022/utils"
	"fmt"
	"sort"
)

func Solve(rawInput []string) (firstSolution string, secondSolution string) {

	parsedInput := splitByEmptyAsInt(rawInput)
	summedInput := sumSections(parsedInput)
	processedInput := sortIntSliceDesc(summedInput)

	firstSolution = fmt.Sprint(processedInput[0])
	secondSolution = fmt.Sprint(utils.SumIntSlice(processedInput[:3]))

	return
}

func sortIntSliceDesc(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)

	sort.Slice(output, func(i, j int) bool {
		return output[i] > output[j]
	})
	return output
}

func sumSections(input [][]int) (output []int) {
	for _, section := range input {
		output = append(output, utils.SumIntSlice(section))
	}
	return
}

func splitByEmptyAsInt(rawInput []string) (output [][]int) {
	var currentSection []int
	for _, v := range rawInput {
		if v != "" {
			currentSection = append(currentSection, utils.ParseInt(v))
		} else {
			output = append(output, currentSection)
			currentSection = nil
		}
	}
	if currentSection != nil {
		output = append(output, currentSection)
	}
	return
}
