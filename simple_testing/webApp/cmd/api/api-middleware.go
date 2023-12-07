package main

import "net/http"

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Acces-Control-Allow-Origin", "http://localhost:8090")

		if r.Method == "OPTIONS" {
			w.Header().Set("Acces-Control-Allow-Credentials", "true")
			w.Header().Set("Acces-Control-Allow-Methods", "GET, POST, PUT,DELETE, PATCH,OPTIONS")
			w.Header().Set("Acces-Control-Allow-Headers", "Accept,Content-Type, X-CSRF-Token, Authorizathion")
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func (app *application) authRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := app.getTokenFromHeaderandVerify(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
