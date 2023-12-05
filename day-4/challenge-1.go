package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Game struct {
	Id            int
	winnerNumbers []int
	yourNumbers   []int
}

func stringSliceToNumSlice(input []string) []int {
	output := make([]int, len(input))

	for i, numStr := range input {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("Could not convert number string to int")
		}

		output[i] = num
	}

	return output
}

func removeEmptyStrings(input []string) []string {
	output := []string{}

	for _, str := range input {
		if str == "" || str == " " {
			continue
		}

		output = append(output, str)
	}

	return output
}

func parseNumList(input string) []int {
	return stringSliceToNumSlice(removeEmptyStrings(strings.Split(input, " ")))
}

func parseInput(input string) []Game {
	lines := strings.Split(input, "\n")
	games := make([]Game, len(lines))

	for i, line := range lines {
		numbersStr := strings.Split(line, ": ")[1]
		numLists := strings.Split(numbersStr, " | ")

		winnerNumbers := parseNumList(numLists[0])
		yourNumbers := parseNumList(numLists[1])

		games[i] = Game{
			Id:            i,
			winnerNumbers: winnerNumbers,
			yourNumbers:   yourNumbers,
		}

	}

	return games
}

func getNumbersMap(numbers []int) map[int]bool {
	numbersMap := make(map[int]bool)
	for _, number := range numbers {
		numbersMap[number] = true
	}

	return numbersMap
}

func Challenge1(input string) {
	games := parseInput(input)

	totalPoints := 0

	for _, game := range games {
		winningNumbersMap := getNumbersMap(game.winnerNumbers)
		matches := 0

		for _, number := range game.yourNumbers {
			if winningNumbersMap[number] {
				matches++
			}
		}

		points := 0

		if matches > 0 {
			points = int(math.Pow(2, (float64)(matches-1)))
			totalPoints += points
		}

	}

	fmt.Printf("%d", totalPoints)
}
