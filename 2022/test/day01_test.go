package test

import (
	"aoc2022/day01"
	"aoc2022/utils"
	"testing"
)

func TestSolverDay01Part1(t *testing.T) {
	expectedFirstSolution := "24000"

	rawInput := utils.ReadFile("day01_testInput.txt")
	firstSolution, _ := day01.Solve(rawInput)
	if firstSolution != expectedFirstSolution {
		t.Fatalf(`Day01 Part1: Result %v did not match expectation %v`, firstSolution, expectedFirstSolution)
	}
}

func TestSolverDay01Part2(t *testing.T) {
	expectedSecondSolution := "45000"

	rawInput := utils.ReadFile("day01_testInput.txt")
	_, secondSolution := day01.Solve(rawInput)

	if secondSolution != expectedSecondSolution {
		t.Fatalf(`Day01 Part1: Result %v did not match expectation %v`, secondSolution, expectedSecondSolution)
	}
}
