package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler) // Show application information
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)     // Create a new movie
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)    // Show the details of a specific movie

	return router
}
