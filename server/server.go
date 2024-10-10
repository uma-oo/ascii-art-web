package asciiArt

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	fn "asciiArt/functions"
)

var Art ascii

type ascii struct {
	Text   string
	Banner string
}

type result struct {
	Text string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("body")
	banner := r.FormValue("banner")

	temp, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	Art = ascii{Text: text, Banner: banner}
	res := handleData()

	temp.Execute(w, res)
}

func handleData() *result {
	// Check for argument errors
	userInput := Art.Banner
	// Open the standard.txt file, read its contents, and create a map of ASCII art

	file, err := os.ReadFile("banners/" + userInput + ".txt")
	if len(file) != 6623 {
		log.Fatal("Error: something wrong with the standard file.")
	}
	if err != nil {
		log.Fatal(err)
	}
	standardArray := strings.Split(string(file[1:len(file)-1]), "\n\n")
	asciiMap := fn.MapBuilder(standardArray)
	// Validate user input
	validUserInput, err := fn.UserInputChecker(userInput)
	if err != nil {
		log.Fatalf("Input Error : %v", err)
	}
	if userInput == "" {
		return &result{""}
	}
	asciiArt := fn.BuildAsciiArt(strings.Split(validUserInput, "\\n"), asciiMap)
	// Check if the user input contains only \n if so we eliminate one of them
	if strings.ReplaceAll(validUserInput, "\\n", "") == "" {
		asciiArt = asciiArt[1:]
	}
	return &result{Text: asciiArt}
}
