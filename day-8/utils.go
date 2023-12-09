package day8

import "strings"

type Point struct {
	name  string
	left  string
	right string
}

func parseInput(input string) (string, []Point) {
	lines := strings.Split(input, "\n")
	directions := lines[0]
	pointsStr := lines[2:]

	points := make([]Point, len(lines)-2)

	for i, pointStr := range pointsStr {
		parts := strings.Split(pointStr, " = ")
		name := parts[0]
		values := strings.Split(parts[1], ", ")
		left := values[0][1:]
		right := values[1][:3]

		points[i] = Point{name, left, right}
	}

	return directions, points
}

func makePointsMap(points []Point) map[string]Point {
	pointsMap := map[string]Point{}

	for _, point := range points {
		pointsMap[point.name] = point
	}

	return pointsMap
}
