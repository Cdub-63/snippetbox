package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getHome)
	mux.HandleFunc("GET /snippet/view/{id}", app.getSnippetView)
	mux.HandleFunc("GET /snippet/create", app.getSnippetCreate)
	mux.HandleFunc("POST /snippet/create", app.postSnippetCreate)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
