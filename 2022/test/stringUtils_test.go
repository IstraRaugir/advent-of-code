package test

import (
	"aoc2022/utils"
	"testing"
)

func TestHalveString(t *testing.T) {
	expectedFirstHalf := "vJrwpWtwJgWr"
	expectedSecondHalf := "hcsFMMfFFhFp"
	fullInput := expectedFirstHalf + expectedSecondHalf
	firstHalf, secondHalf := utils.HalveString(fullInput)
	if firstHalf != expectedFirstHalf || secondHalf != expectedSecondHalf {
		t.Fatalf(`Day03 HalveString(%v): Results %v, %v did not match expectation %v, %v`, fullInput, firstHalf, secondHalf, expectedFirstHalf, expectedSecondHalf)
	}
}
