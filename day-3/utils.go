package day3

import (
	"strconv"
	"strings"
)

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func isNumber(char byte) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func isSymbol(char byte) bool {
	return !isNumber(char) && char != '.'
}

func getNumber(lines []string, line int, startCol int, endCol int) int {
	numStr := ""
	for i := startCol; i <= endCol; i++ {
		numStr = numStr + string(lines[line][i])
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic("Could not parse number")
	}

	return num
}
