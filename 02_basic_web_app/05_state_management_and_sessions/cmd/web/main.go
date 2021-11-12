package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kmrgaurav11235/goweb/state_management_and_sessions/pkg/config"
	"github.com/kmrgaurav11235/goweb/state_management_and_sessions/pkg/handlers"
	"github.com/kmrgaurav11235/goweb/state_management_and_sessions/pkg/render"
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

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
