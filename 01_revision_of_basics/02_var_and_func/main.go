package main

import "fmt"

var myName string // package level variable

func main() {
	fmt.Println("Hello, world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid)

	fmt.Println(sayTwoThings())
}

func saySomething() string {
	return "something"
}

// functions can return more than one things
func sayTwoThings() (string, string) {
	return "something", "else"
}
