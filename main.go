package main

import (
	"log"
	"net/http"

	Handle "asciiartweb/Handles"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle.Home)
	mux.HandleFunc("/AsciiArt", Handle.AsciiArt)
	log.Println("http://localhost:2002/")
	err := http.ListenAndServe(":2002", mux)
	log.Fatal(err)
}
