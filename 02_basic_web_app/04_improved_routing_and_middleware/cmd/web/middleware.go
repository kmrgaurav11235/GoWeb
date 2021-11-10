package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page.")
		next.ServeHTTP(rw, r)
	})
}