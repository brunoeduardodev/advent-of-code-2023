package day3

import "fmt"

func makeChecksMatrix(lines int, cols int) [][]bool {
	matrix := make([][]bool, lines)
	for i := 0; i < lines; i++ {
		matrix[i] = make([]bool, cols)
	}

	return matrix
}

func getNumFromPoint(lines []string, line int, col int, checks [][]bool) int {
	checks[line][col] = true
	startCol := col
	for i := col - 1; i >= 0; i-- {
		if !isNumber(lines[line][i]) {
			break
		}

		checks[line][i] = true
		startCol = i
	}

	endCol := col
	for i := col + 1; i < len(lines[0]); i++ {
		if !isNumber(lines[line][i]) {
			break
		}

		checks[line][i] = true
		endCol = i
	}

	num := getNumber(lines, line, startCol, endCol)
	return num
}

func getAdjacentNumbers(lines []string, line int, col int) []int {
	numbers := []int{}
	checks := makeChecksMatrix(len(lines), len(lines[0]))

	startLine := line
	if startLine > 0 {
		startLine = startLine - 1
	}

	endLine := line
	if endLine < len(lines)-1 {
		endLine = endLine + 1
	}

	startCol := col
	if startCol > 0 {
		startCol = startCol - 1
	}

	endCol := col

	if endCol < len(lines[0])-1 {
		endCol = endCol + 1
	}

	for i := startLine; i <= endLine; i++ {
		for j := startCol; j <= endCol; j++ {
			if checks[i][j] {
				continue
			}

			if !isNumber(lines[i][j]) {
				checks[i][j] = true
				continue
			}

			num := getNumFromPoint(lines, i, j, checks)
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func Challenge2(input string) {
	lines := parseInput(input)
	linesCount := len(lines)
	columnsCount := len(lines[0])
	sumOfGearsRatio := 0

	for i := 0; i < linesCount; i++ {
		for j := 0; j < columnsCount; j++ {
			char := lines[i][j]
			if char != '*' {
				continue
			}

			numbers := getAdjacentNumbers(lines, i, j)
			if len(numbers) > 1 {
				gearRatio := 1
				for _, num := range numbers {
					gearRatio = gearRatio * num
				}

				sumOfGearsRatio += gearRatio
			}
		}
	}

	fmt.Printf("%d", sumOfGearsRatio)
}
