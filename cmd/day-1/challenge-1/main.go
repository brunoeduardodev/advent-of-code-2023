package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brunoeduardodev/advent-of-code/internal"
)

func calculateCalibration(line string) int {
	var firstNumber, lastNumber = -1, -1
	characters := strings.Split(line, "")

	for i := 0; i < len(characters); i++ {
		char := characters[i]
		num, err := strconv.Atoi(char)

		if err != nil {
			continue
		}

		if firstNumber == -1 {
			firstNumber = num
			continue
		}

		lastNumber = num
	}

	if firstNumber == -1 {
		return 0
	}

	if lastNumber == -1 {
		lastNumber = firstNumber
	}

	result := firstNumber*10 + lastNumber

	return result
}

func main() {

	input := internal.GetInput(1, 1)
	lines := strings.Split(input, "\n")

	var sum int = 0
	for i := 0; i < len(lines); i++ {
		res := calculateCalibration(lines[i])
		sum = sum + res
	}

	fmt.Printf("Result=%d", sum)

}
