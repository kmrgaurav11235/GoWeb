package main

import (
	"fmt"
	"net/http"

	"github.com/kmrgaurav11235/goweb/more_on_templates/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // _ ignores the error
	// Run this using `go run cmd/web/*.go`
}
