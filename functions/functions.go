package functions

import (
	"strings"
)

func SplitNewLine(sentence string) []string {
	sentence = strings.ReplaceAll(sentence, "\r\n", "\\n")

	var slice []string
	var word string

	for i := 0; i < len(sentence); i++ {
		if i < len(sentence)-1 && sentence[i] == '\\' && sentence[i+1] == 'n' {

			slice = append(slice, word)
			word = ""
			i++
		} else {
			word += string(sentence[i])
		}
	}

	slice = append(slice, word)

	if strings.ReplaceAll(sentence, "\\n", "") == "" { // "" {"",""}
		slice = slice[1:]
	}

	return slice
}

func CreateMap(list_of_letters []string) map[string][]string {
	m := make(map[string][]string)
	j := 0
	for i := ' '; i < 127; i++ {
		m[string(i)] = strings.Split(list_of_letters[j], "\n")
		j++

	}
	return m
}

func IsPrintable(word string) bool {
	runes := []rune(word)
	result := true
	for i := 0; i < len(runes); i++ {
		if runes[i] > '~' || runes[i] < ' ' {
			result = false
			break
		} else {
			continue
		}
	}
	return result
}

func Print(words []string, m map[string][]string) string {
	new_str := ""
	for _, word := range words {
		if !IsPrintable(word) {
			new_str = "This sentence contains characters out of the range of printable ascii characters"
			break
		}
		if word == "" {
			new_str += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for l := 0; l < len(word); l++ {
				new_str += m[string(word[l])][i]
			}
			new_str += "\n"
		}
	}
	return new_str
}

// func Contains(slice []string, str string) bool {
// 	for _, element := range slice {
// 		if element == str {
// 			return true
// 		}
// 	}
// 	return false
// }
