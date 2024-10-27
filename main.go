package main

import (
	"fmt"
	"log"
	"net/http"

	"asciiArt/server"
)

var Error error


func main() {
	http.HandleFunc("/", server.MainHandler)
	http.HandleFunc("/ascii-art", server.AsciiHandler)

	fmt.Println("Server starting at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
