package day5

import (
	"fmt"
	"math"
	"strconv"
)

func Challenge1(input string) {
	seeds, maps := parseInput(input)

	closestPosition := math.MaxInt

	for _, seedStr := range seeds {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			panic("Could not convert seed")
		}

		currentPosition := seed
		for _, currentMap := range maps {
			validRangeIdx := -1

			for j, mapRange := range currentMap.ranges {

				if currentPosition <= mapRange.endsAt && currentPosition >= mapRange.startAt {
					validRangeIdx = j
				}
			}

			if validRangeIdx != -1 {
				currentPosition += currentMap.ranges[validRangeIdx].delta
			}
		}

		if closestPosition > currentPosition {
			closestPosition = currentPosition
		}

	}

	fmt.Printf("%d", closestPosition)

}
