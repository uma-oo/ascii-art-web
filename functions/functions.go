package functions

import (
	"log"
	"os"
	"strings"
)

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
			new_str += "\r\n"
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

func HandleData(text string, banner string) string {
	var new_file string

	file, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("file: %q", string(file))
	if banner == "thinkertoy" {
		new_file = strings.ReplaceAll(string(file), "\r\n", "\n")
	} else {
		new_file = string(file)
	}

	liste_of_letters := strings.Split(new_file[1:len(new_file)-1], "\n\n")

	m := CreateMap(liste_of_letters)
	words := strings.Split(text, "\r\n")

	return Print(words, m)
}


