package day6

import (
	"strconv"
	"strings"
)

func getLineNumbers(line string) []int {
	numbersStr := strings.Split(line, ":")[1]
	numbers := []int{}
	valuesStr := strings.Split(numbersStr, " ")

	for _, value := range valuesStr {
		if value == "" {
			continue
		}

		num, err := strconv.Atoi(value)
		if err != nil {
			panic("Could not convert number")
		}

		numbers = append(numbers, num)
	}

	return numbers
}

type Race struct {
	time     int
	duration int
}
