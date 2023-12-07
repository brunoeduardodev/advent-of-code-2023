package day6

import (
	"fmt"
	"strings"
)

func parseInput1(input string) []Race {
	lines := strings.Split(input, "\n")
	times := getLineNumbers(lines[0])
	durations := getLineNumbers(lines[1])
	races := make([]Race, len(times))

	for i := range times {
		races[i] = Race{
			time:     times[i],
			duration: durations[i],
		}
	}

	return races
}

func Challenge1(input string) {
	races := parseInput1(input)

	winPossibilities := 1

	for _, race := range races {
		winCount := 0
		for i := 1; i < race.time; i++ {
			if i*(race.time-i) > race.duration {
				winCount++
			}
		}

		winPossibilities *= winCount
	}

	fmt.Printf("%d", winPossibilities)
}
