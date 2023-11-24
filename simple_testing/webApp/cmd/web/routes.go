package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	//registrer middleware
	mux.Use(middleware.Recoverer)
	// registrer routes

	// static assets

	return mux
}
