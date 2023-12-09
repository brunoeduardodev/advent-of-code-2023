package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	numbers := make([][]int, len(lines))

	for i, line := range lines {
		numbersStr := strings.Split(line, " ")
		numbers[i] = make([]int, len(numbersStr))

		for j, numStr := range numbersStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic("Could not convert input")
			}

			numbers[i][j] = num
		}
	}

	return numbers
}

func IsAllZero(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}

	return true
}

func GetNextItemInSequence(sequence []int) int {
	if len(sequence) == 0 {
		panic("Can not get next item on empty sequence")
	} else if len(sequence) == 1 {
		return sequence[0]
	}

	differences := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		differences[i] = sequence[i+1] - sequence[i]
	}

	if IsAllZero(differences) {
		return sequence[0]
	}

	nextDiff := GetNextItemInSequence(differences)
	return sequence[len(sequence)-1] + nextDiff
}

func Challenge1(input string) {
	sequences := parseInput(input)
	sum := 0

	for _, sequence := range sequences {
		nextItem := GetNextItemInSequence(sequence)
		sum += nextItem
	}

	fmt.Printf("%d", sum)
}
