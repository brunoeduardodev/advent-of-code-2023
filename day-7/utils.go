package day7

import (
	"strconv"
	"strings"
)

type GameType uint8

const (
	NO_TYPE GameType = iota
	HIGH_CARD
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

type Game struct {
	cards    []string
	gameType GameType
	bid      int
}

var CardValueMap = map[string]int{
	"*": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

func (gameA Game) isBiggerThan(gameB Game) bool {
	if gameA.gameType != gameB.gameType {
		return gameA.gameType > gameB.gameType
	}

	for i := range gameA.cards {
		if gameA.cards[i] == gameB.cards[i] {
			continue
		}

		return CardValueMap[gameA.cards[i]] > CardValueMap[gameB.cards[i]]
	}

	return true
}

func getGameInfo(cards []string) (distinctCards int, repeatedCards int, jokers int) {
	distinctCards = 0
	repeatedCards = 0
	cardMap := map[string]int{}

	for _, card := range cards {
		if val, ok := cardMap[card]; ok {
			if val == 1 {
				repeatedCards += 2
			} else {
				repeatedCards++
			}

			cardMap[card]++
		} else {
			distinctCards++
			cardMap[card] = 1
		}
	}

	jokers = cardMap["*"]

	return
}

func getGameType(cards []string) GameType {
	distinctCards, repeatedCards, jokers := getGameInfo(cards)
	if distinctCards == 1 {
		return GameType(FIVE_OF_A_KIND)
	} else if distinctCards == 2 && repeatedCards == 4 {
		if jokers == 1 {
			return GameType(FIVE_OF_A_KIND)
		} else if jokers == 4 {
			return GameType(FIVE_OF_A_KIND)
		}

		return GameType(FOUR_OF_A_KIND)
	} else if distinctCards == 2 && repeatedCards == 5 {
		if jokers == 3 {
			return GameType(FIVE_OF_A_KIND)
		} else if jokers == 2 {
			return GameType(FIVE_OF_A_KIND)
		}
		return GameType(FULL_HOUSE)
	} else if distinctCards == 3 && repeatedCards == 3 {
		if jokers == 3 {
			return GameType(FOUR_OF_A_KIND)
		}
		if jokers == 1 {
			return GameType(FOUR_OF_A_KIND)
		}
		return GameType(THREE_OF_A_KIND)
	} else if distinctCards == 3 && repeatedCards == 4 {
		if jokers == 2 {
			return GameType(FOUR_OF_A_KIND)
		} else if jokers == 1 {
			return GameType(FULL_HOUSE)
		}
		return GameType(TWO_PAIR)
	} else if distinctCards == 4 && repeatedCards == 2 {
		if jokers == 2 {
			return GameType(THREE_OF_A_KIND)
		} else if jokers == 1 {
			return GameType(THREE_OF_A_KIND)
		}

		return GameType(ONE_PAIR)
	} else if distinctCards == 5 {
		if jokers == 1 {
			return GameType(ONE_PAIR)
		}
		return GameType(HIGH_CARD)
	}

	return GameType(NO_TYPE)
}

func parseInput(input string, parseJAsJoker bool) []Game {
	lines := strings.Split(input, "\n")
	games := make([]Game, len(lines))

	for i, line := range lines {
		gameValues := strings.Split(line, " ")
		cardsStr := gameValues[0]
		if parseJAsJoker {
			cardsStr = strings.ReplaceAll(cardsStr, "J", "*")
		}
		bidStr := gameValues[1]

		cards := strings.Split(cardsStr, "")
		bid, err := strconv.Atoi(bidStr)

		if err != nil {
			panic("Could not convert bid")
		}

		gameType := getGameType(cards)

		games[i] = Game{cards, gameType, bid}
	}

	return games
}

func quickSortGames(games []Game) []Game {
	if len(games) <= 1 {
		return games
	}

	pivot := games[len(games)-1]

	left := []Game{}
	right := []Game{}

	for i := 0; i < len(games)-1; i++ {
		if games[i].isBiggerThan(pivot) {
			right = append(right, games[i])
		} else {
			left = append(left, games[i])
		}
	}

	left = quickSortGames(left)
	right = quickSortGames(right)

	final := append(left, pivot)
	final = append(final, right...)

	return final
}
