package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/config"
	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	// Set-up the routes
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
