package main

import (
	"fmt"

	day2 "github.com/brunoeduardodev/advent-of-code/cmd/day-2"
	"github.com/brunoeduardodev/advent-of-code/internal"
)

func main() {
	input := internal.GetInput(2, 2)
	games := day2.ParseInput(input)

	powersSum := 0

	for _, game := range games {
		necessaryBallsPerColor := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}

		for _, round := range game.Rounds {
			for _, set := range round.Sets {
				if necessaryBallsPerColor[set.Color] < set.Amount {
					necessaryBallsPerColor[set.Color] = set.Amount
				}
			}
		}

		gamePower := necessaryBallsPerColor["blue"] * necessaryBallsPerColor["green"] * necessaryBallsPerColor["red"]
		powersSum += gamePower
	}

	fmt.Printf("%d", powersSum)
}
