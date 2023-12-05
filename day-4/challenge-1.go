package day4

import (
	"fmt"
	"math"
)

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
