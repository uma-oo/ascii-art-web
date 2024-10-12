package main

import (
	"fmt"
	"log"
	"net/http"

	function "asciiArt/functions"
	se "asciiArt/server"
)

func main() {
	http.HandleFunc("/", se.IndexHandler)
	fmt.Println(function.SplitNewLine("Hi"))
	http.HandleFunc("/ascii-art", se.IndexHandler)
	log.Fatal(http.ListenAndServe(":9008", nil))
}
