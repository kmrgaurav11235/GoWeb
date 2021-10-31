package main

import (
	"errors"
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

func Divide(rw http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(rw, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(rw, "%f divided by %f is %f.", x, y, f)
}

func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// When a function or variable name begins with a capital letter, it means it is accessible outside the given package
// When a function or variable name begins with a small letter, it means the func/var is private

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil) // _ ignores the error
}
