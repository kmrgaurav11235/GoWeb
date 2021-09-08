package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(rw, "Hello World!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Number of Bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil) // _ ignores the error
}
