package main

import (
	"fmt"
	"net/http"
)

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the home page.")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	_, _ = fmt.Fprintf(rw, fmt.Sprintf("This is the about page and 2 + 3 is %d.", sum))

}

// addValues add 2 integers and return the saun.
func addValues(x, y int) int {
	return x + y
}

// When a function or variable name begins with a capital letter, it means it is accessible outside the given package
// When a function or variable name begins with a capital letter, it means the func/var is private

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil) // _ ignores the error
}
