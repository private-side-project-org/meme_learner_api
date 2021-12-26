package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getUserAndMemes(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJson(w, err)
		return
	}

	app.logger.Println("id is", id)

	user, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, user, "user")

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getRandomMeme(w http.ResponseWriter, r *http.Request) {
	scrapedMeme, err := app.models.DB.Scraping()
	if err != nil {
		app.logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, scrapedMeme, "scrapedMeme")

	if err != nil {
		app.logger.Println(err)
	}
}
