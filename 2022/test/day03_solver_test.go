package test

import (
	"aoc2022/day03"
	"aoc2022/utils"
	"fmt"
	"testing"
)

func TestCharToPriority(t *testing.T) {
	fullAlphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, v := range []rune(fullAlphabet) {
		output := day03.CharToPriority(v)
		expectation := i + 1
		if expectation != int(output) {
			fmt.Println(output, expectation)
			t.Fatalf(`Day03 CharToPriority(%v): Result %v did not match expectation %v`, string(v), output, expectation)
		}
	}
}

func TestSolverDay03Part1(t *testing.T) {
	expectedFirstSolution := "157"

	rawInput := utils.ReadFile("day03_testInput.txt")
	firstSolution, _ := day03.Solve(rawInput)
	if firstSolution != expectedFirstSolution {
		t.Fatalf(`Day03 Part1: Result %v did not match expectation %v`, firstSolution, expectedFirstSolution)
	}
}

func TestSolverDay03Part2(t *testing.T) {
	expectedSecondSolution := "70"

	rawInput := utils.ReadFile("day03_testInput.txt")
	_, secondSolution := day03.Solve(rawInput)

	if secondSolution != expectedSecondSolution {
		t.Fatalf(`Day03 Part1: Result %v did not match expectation %v`, secondSolution, expectedSecondSolution)
	}
}
