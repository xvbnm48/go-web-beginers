package main

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
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
	mux.Use(a.LoadSession)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	// routes for the application.
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		a.session.Put(r.Context(), "test", "sakura endo")
		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Fatal(err)
		}
	}) // GET /

	mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		vars := make(jet.VarMap)
		vars.Set("test", a.session.GetString(r.Context(), "test"))
		err := a.render(w, r, "index", vars)
		if err != nil {
			log.Fatal(err)
		}

	}) // GET /comments

	return mux
}
