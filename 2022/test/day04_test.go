package test

import (
	"aoc2022/day04"
	"aoc2022/utils"
	"testing"
)

func TestSolverDay04Part1(t *testing.T) {
	expectedFirstSolution := "2"

	rawInput := utils.ReadFile("day04_testInput.txt")
	firstSolution, _ := day04.Solve(rawInput)
	if firstSolution != expectedFirstSolution {
		t.Fatalf(`Day04 Part1: Result %v did not match expectation %v`, firstSolution, expectedFirstSolution)
	}
}

func TestSolverDay04Part2(t *testing.T) {
	expectedSecondSolution := "4"

	rawInput := utils.ReadFile("day04_testInput.txt")
	_, secondSolution := day04.Solve(rawInput)

	if secondSolution != expectedSecondSolution {
		t.Fatalf(`Day04 Part1: Result %v did not match expectation %v`, secondSolution, expectedSecondSolution)
	}
}
