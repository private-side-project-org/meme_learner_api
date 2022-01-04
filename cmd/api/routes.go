package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// wrapper function to get context of request and call next handler
func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// get context of request with parameters
		ctx := context.WithValue(r.Context(), "params", ps)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// router to exec action when request happen on each path
func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	// sign up
	router.HandlerFunc(http.MethodPost, "/v1/session", app.SignUp)

	// sign in
	router.HandlerFunc(http.MethodPost, "/v1/login", app.Signin)

	// exec status handler when GET request on /status
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// exec getUserAndMemes when GET request on /v1/user/:id
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.getUserAndMemes)

	// exec getRnadomMeme with Auth token!!
	router.GET("/v1/memes/random", app.wrap(secure.ThenFunc(app.getRandomMeme)))

	return app.enableCORS(router)
}
