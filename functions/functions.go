package functions

import (
	"errors"
	"strings"
)

// MapBuilder takes a slice of strings and returns a map where the keys are
// printable ASCII characters and the values are the corresponding ASCII art strings.
func MapBuilder(data []string) map[rune][]string {
	mapHolder := map[rune][]string{}
	indexCounter := 0
	for i := ' '; i <= '~'; i++ {
		mapHolder[i] = strings.Split(data[indexCounter], "\n")
		indexCounter++
	}
	return mapHolder
}

func UserInputChecker(input string) (string, error) {
	for _, i := range input {
		if i < 32 || i > 126 {
			return "", errors.New("found a character outside the range of printable ascii characters")
		}
	}
	return input, nil
}

func BuildAsciiArt(input []string, asciiMap map[rune][]string) string {
	result := ""
	for _, i := range input {
		if i == "" {
			result += "\n"
			continue
		}
		for j := 0; j < 8; j++ {
			for _, k := range i {
				result += asciiMap[k][j]
			}
			result += "\n"
		}
	}
	return result
}
