package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"memelearner_be/models"
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
	"golang.org/x/crypto/bcrypt"
)

// placeholder user struct
var validUser = models.User{
	ID:       10,
	UserName: "testuser",
	Password: "$2a$12$8rTJOPyLkiv5sqGIe32O8eANHGjGToxtiX78pogzKtWwYMqd.7uvm",
}

// request body receiver
type Credentials struct {
	UserName string `json:"id"`
	Password string `json:"password"`
}

func (app *application) Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		app.errorJson(w, errors.New("Unauthorized"))
		return
	}

	// password from DB(temporary used placeholder user struct)
	hashedPassword := validUser.Password

	// password from client
	clientPassword := creds.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(clientPassword))

	if err != nil {
		fmt.Println("error when crypting")
		fmt.Printf("why crypting %v", &creds)
		app.errorJson(w, errors.New("Unauthorized"))
		return
	}

	var claims jwt.Claims
	claims.Subject = fmt.Sprint(validUser.ID)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "mydomain.com"
	claims.Audiences = []string{"mydomain.com"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))

	if err != nil {
		app.errorJson(w, errors.New("error signing"))
		return
	}

	err = app.writeJSON(w, http.StatusOK, jwtBytes, "response")
}
