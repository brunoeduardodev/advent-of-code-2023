package day8

import (
	"fmt"
)

func Challenge1(input string) {
	directions, points := parseInput(input)
	pointsMap := makePointsMap(points)
	currentPoint := "AAA"
	steps := 0

	for i := 0; ; i++ {
		if currentPoint == "ZZZ" {
			fmt.Printf("%d", steps)
			return
		}

		direction := directions[i%len(directions)]
		if direction == 'L' {
			currentPoint = pointsMap[currentPoint].left
		} else {
			currentPoint = pointsMap[currentPoint].right
		}

		steps++
	}
}
