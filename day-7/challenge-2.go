package day7

import "fmt"

func Challenge2(input string) {
	games := parseInput(input, true)
	sortedGames := quickSortGames(games)

	total := 0

	for i, game := range sortedGames {
		rank := i + 1
		total += rank * game.bid
	}

	fmt.Printf("%d", total)

}
