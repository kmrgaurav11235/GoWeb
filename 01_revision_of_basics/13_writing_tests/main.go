package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 0)

	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Result of division is", result)
	}
}

func divide(dividend, divisor float32) (float32, error) {
	var result float32

	if divisor == 0 {
		return result, errors.New("division by 0 is not allowed") // Throw errors like this
	} else {
		result = dividend / divisor
		return result, nil
	}
}
