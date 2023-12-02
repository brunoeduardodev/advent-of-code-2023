package day2

import (
	"strconv"
	"strings"
)

type Set struct {
	Color  string
	Amount int
}

type Round struct {
	Sets []Set
}

type Game struct {
	Id     int
	Rounds []Round
}

func ParseSet(rawSet string) Set {
	data := strings.Split(rawSet, " ")
	amount, err := strconv.Atoi(data[0])

	if err != nil {
		panic("Could not convert amount to int")
	}

	color := data[1]

	set := Set{
		Color:  color,
		Amount: amount,
	}

	return set
}

func ParseRound(rawRound string) Round {
	rawSets := strings.Split(rawRound, ", ")
	sets := make([]Set, len(rawSets))

	for i, rawSet := range rawSets {
		sets[i] = ParseSet(rawSet)
	}

	round := Round{Sets: sets}
	return round
}

func ParseGame(index int, raw string) Game {
	roundsText := strings.Split(raw, ": ")[1]
	rawRounds := strings.Split(roundsText, "; ")

	rounds := make([]Round, len(rawRounds))

	for i, round := range rawRounds {
		rounds[i] = ParseRound(round)
	}

	game := Game{
		Id:     index + 1,
		Rounds: rounds,
	}

	return game
}

func ParseInput(input string) []Game {
	rawGames := strings.Split(input, "\n")
	games := make([]Game, len(rawGames))

	for i, rawGame := range rawGames {
		games[i] = ParseGame(i, rawGame)
	}

	return games
}
