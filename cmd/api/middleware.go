package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

// middleware to inject cors accpets
func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// adding `Access-Control-Allow-Origin: *` which accepts all
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		// if request pre-flight(which is OPTION), return OK status(which means, accept everything)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get(("Authorization"))

		// set for the case handle when login anonymous user
		if authHeader == "" {
			// could set an anonymous user
		}

		// array of headers
		headerParts := strings.Split(authHeader, " ")

		if len(headerParts) != 2 {
			app.errorJson(w, errors.New("invalid auth header"))
			return
		}

		if headerParts[0] != "Bearer" {
			app.errorJson(w, errors.New("unauthorized - no bearer"))
			return
		}

		token := headerParts[1]

		// HMAC check token and secret, then return claims
		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))

		// HMAC error
		if err != nil {
			app.errorJson(w, errors.New("unauthorized - failed HMACCheck"), http.StatusForbidden)
			return
		}

		/* claims check*/
		// token expire
		if !claims.Valid(time.Now()) {
			app.errorJson(w, errors.New("unauthorized - token expired"), http.StatusForbidden)
			return
		}

		// invalid audiences
		if !claims.AcceptAudience("mydomain.com") {
			app.errorJson(w, errors.New("unauthorized - invalid audience"), http.StatusForbidden)
			return
		}

		// invlid issuer
		if claims.Issuer != "mydomain.com" {
			app.errorJson(w, errors.New("unauthorized - invalid issuer"), http.StatusForbidden)
			return
		}

		/* finish claims check */

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)

		if err != nil {
			app.errorJson(w, errors.New("unauthorized, something wrong on claims.Subject"), http.StatusForbidden)
			return
		}

		log.Println("Valid user", userID)

		next.ServeHTTP(w, r)
	})
}
