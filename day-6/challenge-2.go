package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumberFromNumbersSlice(numbers []int) int {
	numStr := ""

	for _, num := range numbers {
		numStr = fmt.Sprintf("%s%d", numStr, num)
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic("could not convert number")
	}

	return num
}

func parseInput2(input string) Race {
	lines := strings.Split(input, "\n")
	times := getLineNumbers(lines[0])

	durations := getLineNumbers(lines[1])

	time := getNumberFromNumbersSlice(times)
	duration := getNumberFromNumbersSlice(durations)

	return Race{
		time:     time,
		duration: duration,
	}

}

func Challenge2(input string) {
	race := parseInput2(input)

	winCount := 0
	for i := 1; i < race.time; i++ {
		if i*(race.time-i) > race.duration {
			winCount++
		}
	}

	fmt.Printf("%d", winCount)
}
