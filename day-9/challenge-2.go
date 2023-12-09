package day9

import (
	"fmt"
)

func reverse[T any](list []T) []T {
	reversed := make([]T, len(list))

	for i := range list {
		reversed[i] = list[len(list)-i-1]
	}

	return reversed
}

func Challenge2(input string) {
	sequences := parseInput(input)
	sum := 0

	for _, sequence := range sequences {
		nextItem := GetNextItemInSequence(reverse(sequence))
		sum += nextItem
	}

	fmt.Printf("%d", sum)
}
