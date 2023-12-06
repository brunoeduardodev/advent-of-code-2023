package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type SeedRange struct {
	startsAt int
	endsAt   int
}

func reverseArray[T any](arr []T) []T {
	reverse := make([]T, len(arr))
	for i := 0; i < len(arr); i++ {
		reverse[i] = arr[len(arr)-1-i]
	}

	return reverse
}

func parseInput2(input string) ([]string, []Map) {
	blocks := strings.Split(input, "\n\n")
	seeds := strings.Split(strings.Split(blocks[0], ": ")[1], " ")

	maps := make([]Map, 7)
	for i := 1; i < len(blocks); i++ {
		blockLines := strings.Split(blocks[i], "\n")
		blockMap := Map{}
		blockMap.ranges = make([]MapRange, len(blockLines)-1)

		for j := 1; j < len(blockLines); j++ {
			lineValues := strings.Split(blockLines[j], " ")

			sourceRangeStart, err := strconv.Atoi(lineValues[0])
			if err != nil {
				panic("Could not convert destination range")
			}

			destinationRangeStart, err := strconv.Atoi(lineValues[1])
			if err != nil {
				panic("Could not convert source range")
			}

			rangeLength, err := strconv.Atoi(lineValues[2])
			if err != nil {
				panic("Could not convert range length")
			}

			blockMap.ranges[j-1] = MapRange{}

			blockMap.ranges[j-1].startAt = sourceRangeStart
			blockMap.ranges[j-1].endsAt = sourceRangeStart + rangeLength
			blockMap.ranges[j-1].delta = destinationRangeStart - sourceRangeStart
		}

		maps[i-1] = blockMap
	}

	return seeds, maps
}

func getSeeds(seedValues []string) []SeedRange {
	seeds := make([]SeedRange, len(seedValues)/2)
	for i := 0; i < len(seedValues)/2; i++ {
		startRange, err := strconv.Atoi(seedValues[i*2])
		if err != nil {
			panic("Could not convert start range")
		}

		rangeLength, err := strconv.Atoi(seedValues[i*2+1])
		if err != nil {
			panic("Could not convert start range")
		}

		seeds[i] = SeedRange{
			startsAt: startRange,
			endsAt:   startRange + rangeLength,
		}

	}

	return seeds
}

func Challenge2(input string) {
	seedsValues, orderedMaps := parseInput2(input)
	seeds := getSeeds(seedsValues)
	maps := reverseArray(orderedMaps)

	i := 0
	for {
		currentPosition := i
		for _, currentMap := range maps {
			validRangeIdx := -1

			for j, mapRange := range currentMap.ranges {
				if currentPosition < mapRange.endsAt && currentPosition >= mapRange.startAt {
					validRangeIdx = j
				}
			}

			if validRangeIdx != -1 {
				currentPosition += currentMap.ranges[validRangeIdx].delta
			}

			if currentPosition < 0 {
				break
			}
		}

		if currentPosition < 0 {
			i++
			continue
		}

		for _, seedRange := range seeds {
			if seedRange.startsAt <= currentPosition && seedRange.endsAt > currentPosition {
				fmt.Printf("%d", i)
				return
			}
		}

		i++
	}

}
