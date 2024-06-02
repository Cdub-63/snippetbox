package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Get Methods
	mux.HandleFunc("GET /{$}", getHome)
	mux.HandleFunc("GET /snippet/view/{id}", getSnippetView)
	mux.HandleFunc("GET /snippet/create", getSnippetCreate)

	// Post Methods
	mux.HandleFunc("POST /snippet/create", postSnippetCreate)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
