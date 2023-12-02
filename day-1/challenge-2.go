package day1

import (
	"fmt"
	"strconv"
	"strings"
)

var spelledNumbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func isValidBuffer(buffer string) bool {
	for i := 0; i < len(spelledNumbers); i++ {
		if strings.HasPrefix(spelledNumbers[i], buffer) {
			return true
		}
	}

	return false
}

func getNumber(spelled string) int {
	for i := 0; i < len(spelledNumbers); i++ {
		if spelled == spelledNumbers[i] {
			return i + 1
		}
	}

	return -1
}

func clearBuffer(buffer string) string {
	for i := 0; i < len(buffer); i++ {
		if isValidBuffer(buffer[i:]) {
			return buffer[i:]
		}
	}

	return ""
}

func day2CalculateCalibration(line string) int {

	var firstNumber, lastNumber = -1, -1
	characters := strings.Split(line, "")

	wordBuffer := ""

	for i := 0; i < len(characters); i++ {
		char := characters[i]
		num, err := strconv.Atoi(char)

		if err != nil {
			wordBuffer += char

			isValid := isValidBuffer(wordBuffer)
			if !isValid {
				wordBuffer = clearBuffer(wordBuffer)
				continue
			}

			convertedNum := getNumber(wordBuffer)
			if convertedNum == -1 {
				continue
			}

			wordBuffer = clearBuffer(wordBuffer[1:])
			num = convertedNum
		} else {
			wordBuffer = ""
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

func Challenge2(input string) {
	lines := strings.Split(input, "\n")

	var sum int = 0
	for i := 0; i < len(lines); i++ {
		res := day2CalculateCalibration(lines[i])
		sum = sum + res
	}

	fmt.Print(sum)
}
