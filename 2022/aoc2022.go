package main

import (
	"aoc2022/day01"
	"aoc2022/utils"
)

func main() {
	rawInput := utils.ReadFile("day01/input.txt")
	a, b := day01.Solve(rawInput)
	utils.PrintSolutions("01", a, b)
}
