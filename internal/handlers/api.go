package handlers

import (
	"github.com/Fox658/api_go_exercise/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		// Middleware for  /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)

	})

}
