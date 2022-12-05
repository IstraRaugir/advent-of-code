package utils

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) (output []string) {
	file, err := os.Open(fileName)
	PanicIfErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	PanicIfErr(scanner.Err())
	return
}
