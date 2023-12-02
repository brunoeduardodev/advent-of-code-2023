package helpers

import (
	"fmt"
	"os"
)

func GetInput(day int, challenge int) string {
	pwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get cwd")
	}

	path := fmt.Sprintf("%s/inputs/day-%d/challenge-%d.txt", pwd, day, challenge)
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Unable to read file")
	}

	return string(data)
}

type Day map[int]func(string)
