package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// router to exec action when request happen on each path
func (app *application) routes() http.Handler {
	router := httprouter.New()

	// exec status handler when GET request on /status
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// exec getUserAndMemes when GET request on /v1/user/:id
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.getUserAndMemes)

	// exec getRandomMeme when GET request on /v1/memes/random
	router.HandlerFunc(http.MethodGet, "/v1/memes/random", app.getRandomMeme)

	return app.enableCORS(router)
}
