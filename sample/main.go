package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/api", func(r chi.Router) {
		sampleRouter(r)
	})
	http.ListenAndServe(":3000", r)
}
