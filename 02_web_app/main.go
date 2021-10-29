package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the home page.")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	_, _ = fmt.Fprintf(rw, "This is the about page and 2 + 3 is %d.", sum)

}

// addValues add 2 integers and return the saun.
func addValues(x, y int) int {
	return x + y
}

// When a function or variable name begins with a capital letter, it means it is accessible outside the given package
// When a function or variable name begins with a small letter, it means the func/var is private

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // _ ignores the error
}
