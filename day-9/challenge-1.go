package day9

import (
	"fmt"
)

func Challenge1(input string) {
	sequences := parseInput(input)
	sum := 0

	for _, sequence := range sequences {
		nextItem := GetNextItemInSequence(sequence)
		sum += nextItem
	}

	fmt.Printf("%d", sum)
}
