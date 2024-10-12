package asciiArt

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	function "asciiArt/functions"
)

type data struct {
	Input  string
	Banner string
	Output string
}

type input struct {
	Text   string
	Banner string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 4096) 

	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, &data{})
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() Error: %v", err)
			return
		}
		text := r.FormValue("body")
		banner := r.FormValue("banner")

		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		a := &input{Text: text, Banner: banner}
		if (input{}) == *a {
			temp.Execute(w, &a)
		} else {
			temp.Execute(w, &data{Input: a.Text, Banner: a.Banner, Output: HandleData(a.Text, a.Banner)})
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
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

	m := function.CreateMap(liste_of_letters)
	words := function.SplitNewLine(text)

	return function.Print(words, m)
}
