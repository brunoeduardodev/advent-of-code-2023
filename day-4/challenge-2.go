package day4

import (
	"fmt"
)

func Challenge2(input string) {
	games := parseInput(input)

	copiesCount := make([]int, len(games))
	fillArrayWithOne(copiesCount)

	for i, game := range games {
		matches := getMatches(game)

		for j := 1; j <= matches; j++ {
			copiesCount[i+j] += 1 * copiesCount[i]
		}
	}

	total := 0
	for _, count := range copiesCount {
		total += count
	}

	fmt.Printf("%d", total)
}
