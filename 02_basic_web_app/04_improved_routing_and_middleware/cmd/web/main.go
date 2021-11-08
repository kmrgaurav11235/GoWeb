package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/config"
	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/handlers"
	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // _ ignores the error
	// Run this using `go run cmd/web/*.go`
}
