package day3

import (
	"fmt"
)

func isAdjacentToSymbols(lines []string, line int, numStartColumn int, numEndColumn int) bool {
	startColumn := numStartColumn
	if startColumn > 0 {
		startColumn = startColumn - 1
	}

	endColumn := numEndColumn
	if endColumn < len(lines[0])-1 {
		endColumn = endColumn + 1
	}

	startLine := line
	if startLine > 0 {
		startLine = startLine - 1
	}

	endLine := line
	if endLine < len(lines)-1 {
		endLine = endLine + 1
	}

	for i := startLine; i <= endLine; i++ {
		for j := startColumn; j <= endColumn; j++ {
			if isSymbol(lines[i][j]) {
				return true
			}
		}
	}

	return false
}

func Challenge1(input string) {
	lines := parseInput(input)

	linesCount := len(lines)
	columnsCount := len(lines[0])

	partNumbersSum := 0

	for i := 0; i < linesCount; i++ {
		for j := 0; j < columnsCount; j++ {
			char := lines[i][j]
			if isNumber(char) {
				startColumn := j
				endColumn := j

				for k := j; k < columnsCount; k++ {
					possibleNumberChar := lines[i][k]
					if !isNumber(possibleNumberChar) {
						break
					}

					// fmt.Printf("Number found %s\n", string(possibleNumberChar))

					endColumn = k
					j = k
				}

				isPartNumber := isAdjacentToSymbols(lines, i, startColumn, endColumn)
				if isPartNumber {
					num := getNumber(lines, i, startColumn, endColumn)
					partNumbersSum += num
				}
			}
		}
	}

	fmt.Printf("%d", partNumbersSum)
}
