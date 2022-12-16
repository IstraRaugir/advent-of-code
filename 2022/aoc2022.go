package main

import (
	"aoc2022/day01"
	"aoc2022/day02"
	"aoc2022/day03"
	"aoc2022/day04"
	"aoc2022/utils"
	"flag"
	"fmt"
)

func main() {
	dayToRun := flag.Int("day", 0, "which day to run")
	flag.Parse()

	if *dayToRun != 0 {
		runForDay(*dayToRun)
	} else {
		for i := 1; i <= 25; i++ {
			runForDay(i)
		}
	}
}

func runForDay(day int) {

	var rawInput []string
	var firstSolution, secondSolution string
	dayString := mapDayToString(day)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Day", dayString, "encountered problem:", r)
		}
	}()

	rawInput = utils.ReadFile("day" + dayString + "/input.txt")
	firstSolution, secondSolution = solveForDay(day, rawInput)
	utils.PrintSolutions(dayString, firstSolution, secondSolution)
}

func solveForDay(day int, rawInput []string) (string, string) {
	switch day {
	case 1:
		return day01.Solve(rawInput)
	case 2:
		return day02.Solve(rawInput)
	case 3:
		return day03.Solve(rawInput)
	case 4:
		return day04.Solve(rawInput)
	default:
		panic("unimplemented")
	}
}

func mapDayToString(day int) string {
	dayString := ""
	if day < 10 {
		dayString += "0"
	}
	dayString += fmt.Sprint(day)
	return dayString
}
