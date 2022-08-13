package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes: define the routes for the application.
func (a *application) routes() http.Handler {
	// Create a new router.
	mux := chi.NewRouter()

	// Add a middleware that will log the start and end of each request.
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	// routes for the application.
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(a.appName))
	}) // GET /

	mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("comments"))
	}) // GET /comments

	return mux
}
