package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // _ ignores the error
}
