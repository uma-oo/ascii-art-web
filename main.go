package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"asciiArt/server"
)

var Error error

var Err_Passed string

func init() {
	templates, Error := template.ParseGlob("templates/*.html")
	if Error != nil {
		Err_Passed = fmt.Sprintf("%s", Error)
	}
	server.Init(templates)
	Err_Passed=""
}

func main() {
	http.HandleFunc("/",server.MainHandler)
	http.HandleFunc("/ascii-art", server.AsciiHandler)

	fmt.Println("Server starting at http://localhost:8080")
	
	log.Fatal(http.ListenAndServe(":8080",nil))

}
