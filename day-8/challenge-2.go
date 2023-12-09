package day8

import "fmt"

func GetGhostEntrypoints(points []Point) []string {
	ghostEntrypoints := []string{}

	for _, point := range points {
		if point.name[2] == 'A' {
			ghostEntrypoints = append(ghostEntrypoints, point.name)
		}
	}

	return ghostEntrypoints
}

func GetLCD(numbers []int) int {
	divisors := []int{}
	currentDivisor := 2

	currentNumbers := numbers

	for {
		areAllNumbersDivided := true
		for _, num := range currentNumbers {
			if num != 1 {
				areAllNumbersDivided = false
			}
		}

		if areAllNumbersDivided {
			leastCommonDivisor := 1
			for _, divisor := range divisors {
				leastCommonDivisor *= divisor
			}

			return leastCommonDivisor
		}

		allSkipped := true
		for i, num := range currentNumbers {
			if num%currentDivisor == 0 {
				currentNumbers[i] = currentNumbers[i] / currentDivisor
				allSkipped = false
			}
		}

		if allSkipped {
			currentDivisor += 1
		} else {
			divisors = append(divisors, currentDivisor)
		}
	}
}

func Challenge2(input string) {
	directions, points := parseInput(input)
	pointsMap := makePointsMap(points)
	currentPoints := GetGhostEntrypoints(points)
	currentPointsCycles := make([][]int, len(currentPoints))
	for i := range currentPointsCycles {
		currentPointsCycles[i] = []int{}
	}

	steps := 0

	for i := 0; ; i++ {
		direction := directions[i%len(directions)]
		areAllFinished := true

		for j := range currentPoints {
			if currentPoints[j][2] != 'Z' {
				areAllFinished = false
				break
			}
		}

		if areAllFinished {
			fmt.Printf("%d", steps)
			return
		}

		for j := range currentPoints {
			if currentPoints[j][2] == 'Z' {
				isCycleSaved := false
				for k := range currentPointsCycles[j] {
					if steps%currentPointsCycles[j][k] == 0 {
						isCycleSaved = true
					}
				}

				if !isCycleSaved {
					currentPointsCycles[j] = append(currentPointsCycles[j], steps)
				}
			}
		}

		for j := range currentPoints {
			if direction == 'L' {
				currentPoints[j] = pointsMap[currentPoints[j]].left
			} else {
				currentPoints[j] = pointsMap[currentPoints[j]].right
			}

		}

		allHaveCycles := true
		for j := range currentPointsCycles {
			if len(currentPointsCycles[j]) == 0 {
				allHaveCycles = false
			}
		}

		flatCycles := []int{}
		for _, cycle := range currentPointsCycles {
			flatCycles = append(flatCycles, cycle...)
		}

		if allHaveCycles {
			fmt.Printf("%v\n", GetLCD(flatCycles))
			return
		}

		steps++
	}
}
