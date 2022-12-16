package test

import (
	"aoc2022/day02"
	"aoc2022/utils"
	"testing"
)

func TestSolverDay02Part1(t *testing.T) {
	expectedFirstSolution := "15"

	rawInput := utils.ReadFile("day02_testInput.txt")
	firstSolution, _ := day02.Solve(rawInput)
	if firstSolution != expectedFirstSolution {
		t.Fatalf(`Day02 Part1: Result %v did not match expectation %v`, firstSolution, expectedFirstSolution)
	}
}

func TestSolverDay02Part2(t *testing.T) {
	expectedSecondSolution := "12"

	rawInput := utils.ReadFile("day02_testInput.txt")
	_, secondSolution := day02.Solve(rawInput)

	if secondSolution != expectedSecondSolution {
		t.Fatalf(`Day02 Part1: Result %v did not match expectation %v`, secondSolution, expectedSecondSolution)
	}
}
