package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.getUserAndMemes)
	router.HandlerFunc(http.MethodGet, "/v1/memes/random", app.getRandomMeme)

	return app.enableCORS(router)
}