package day2

import (
	"fmt"
)

func Challenge2(input string) {
	games := ParseInput(input)

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
