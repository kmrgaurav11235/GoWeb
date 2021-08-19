package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

const numPool = 1000

func CalculateValue(c chan int) {
	randomNum := RandomNumber(numPool)
	c <- randomNum
}

func main() {
	// Creating a channel is similar to creating maps - use make keyword
	intChan := make(chan int) // channel of type string
	defer close(intChan)      //defer -> Execute whatever comes after me as soon as the current function is done

	go CalculateValue(intChan) // This will start a new Go Routine to run this function.

	num := <-intChan

	fmt.Println("Got", num, "from the channel")
}
