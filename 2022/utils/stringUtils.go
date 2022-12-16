package utils

import "strconv"

func ParseInt(v string) int {
	intValue, err := strconv.Atoi(v)
	PanicIfErr(err)
	return intValue
}

func HalveString(input string) (string, string) {
	cutPoint := len(input) / 2
	return input[:cutPoint], input[cutPoint:]
}
