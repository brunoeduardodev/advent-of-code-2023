package day2

import (
	"fmt"
)

func Challenge1(input string) {
	games := ParseInput(input)

	maxBallsPerColor := map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}

	validGameIdsSum := 0

	for i, game := range games {
		isPossible := true

		for _, round := range game.Rounds {
			for _, set := range round.Sets {
				if maxBallsPerColor[set.Color] < set.Amount {
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
