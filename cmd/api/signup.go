package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"memelearner_be/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *application) SignUp(w http.ResponseWriter, r *http.Request) {
	var uc models.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&uc)

	if err != nil {
		app.errorJson(w, errors.New("unauthorized - failed to decode"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(uc.Password), 14)

	if err != nil {
		app.errorJson(w, errors.New("unauthorized - failed to hash"))
		return
	}

	uc.Password = string(hashedPassword)

	uCreds := app.models.DB.CreateUser(&uc)

	fmt.Printf(`value is %v`, uCreds)

	app.writeJSON(w, http.StatusOK, uCreds, "sign up")
}
