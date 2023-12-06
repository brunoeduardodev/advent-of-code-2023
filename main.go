package main

import (
	"fmt"
	"os"
	"strconv"

	day1 "github.com/brunoeduardodev/advent-of-code/day-1"
	day2 "github.com/brunoeduardodev/advent-of-code/day-2"
	day3 "github.com/brunoeduardodev/advent-of-code/day-3"
	day4 "github.com/brunoeduardodev/advent-of-code/day-4"
	day5 "github.com/brunoeduardodev/advent-of-code/day-5"
	"github.com/brunoeduardodev/advent-of-code/helpers"
)

func parseArgs() (int, int, error) {
	args := os.Args[1:]

	if len(args) != 2 {
		return -1, -1, fmt.Errorf("invalid args count! Please input the day and the challenge")
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, -1, fmt.Errorf("day needs to be an integer")
	}

	_, ok := daysHandlers[day]

	if !ok {
		return -1, -1, fmt.Errorf(fmt.Sprintf("could not find day %d", day))
	}

	challenge, err := strconv.Atoi(args[1])

	if err != nil {
		return -1, -1, fmt.Errorf("challenge needs to be an integer")
	}

	if challenge != 1 && challenge != 2 {
		return -1, -1, fmt.Errorf("invalid challenge number it needs to be either 1 or 2")
	}

	return day, challenge, nil
}

var daysHandlers = map[int]helpers.Day{
	1: day1.Day1,
	2: day2.Day2,
	3: day3.Day3,
	4: day4.Day4,
	5: day5.Day5,
}

func main() {
	day, challenge, err := parseArgs()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	handler := daysHandlers[day]

	input := helpers.GetInput(day, challenge)
	handler[challenge](input)
}
