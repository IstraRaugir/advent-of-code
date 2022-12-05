package main

import (
	"aoc2022/day01"
	"aoc2022/day02"
	"aoc2022/utils"
)

func main() {
	var rawInput []string
	var firstSolution, secondSolution string

	rawInput = utils.ReadFile("day01/input.txt")
	firstSolution, secondSolution = day01.Solve(rawInput)
	utils.PrintSolutions("01", firstSolution, secondSolution)

	rawInput = utils.ReadFile("day02/input.txt")
	firstSolution, secondSolution = day02.Solve(rawInput)
	utils.PrintSolutions("02", firstSolution, secondSolution)
}
