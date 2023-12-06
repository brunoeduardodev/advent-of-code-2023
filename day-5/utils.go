package day5

import (
	"strconv"
	"strings"
)

type MapRange struct {
	startAt int
	endsAt  int
	delta   int
}

type Map struct {
	ranges []MapRange
}

func parseInput(input string) ([]string, []Map) {
	blocks := strings.Split(input, "\n\n")
	seeds := strings.Split(strings.Split(blocks[0], ": ")[1], " ")

	maps := make([]Map, 7)
	for i := 1; i < len(blocks); i++ {
		blockLines := strings.Split(blocks[i], "\n")
		blockMap := Map{}
		blockMap.ranges = make([]MapRange, len(blockLines)-1)

		for j := 1; j < len(blockLines); j++ {
			lineValues := strings.Split(blockLines[j], " ")

			destinationRangeStart, err := strconv.Atoi(lineValues[0])
			if err != nil {
				panic("Could not convert destination range")
			}

			sourceRangeStart, err := strconv.Atoi(lineValues[1])
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
