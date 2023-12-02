package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brunoeduardodev/advent-of-code/internal"
)

func main() {
	input := internal.GetInput(2, 1)
	games := strings.Split(input, "\n")

	maxBallsPerColor := map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}

	validGameIdsSum := 0

	for i, game := range games {
		roundsText := strings.Split(game, ": ")[1]
		rounds := strings.Split(roundsText, "; ")

		isPossible := true
		for _, round := range rounds {
			sets := strings.Split(round, ", ")

			for _, set := range sets {
				data := strings.Split(set, " ")
				amount, err := strconv.Atoi(data[0])

				if err != nil {
					panic("Could not convert amount to int")
				}

				color := data[1]

				if maxBallsPerColor[color] < amount {
					isPossible = false
					continue
				}
			}

			if !isPossible {
				continue
			}
		}

		if isPossible {
			validGameIdsSum += i + 1
		}
	}

	fmt.Printf("%d", validGameIdsSum)
}
