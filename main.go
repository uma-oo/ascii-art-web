package main

import (
	"log"
	"net/http"

	se "asciiArt/server"
)

func main() {
	http.HandleFunc("/hi", se.IndexHandler)
	log.Fatal(http.ListenAndServe(":9008", nil))
}
