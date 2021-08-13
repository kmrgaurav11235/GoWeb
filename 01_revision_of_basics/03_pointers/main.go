package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) // &variable = Give me the memory address of the value that this variable is pointing at.

	log.Println("After function call, myString is set to", myString)
}

// Below *string is a type description. It means we are working with a "Pointer to a string".
func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
	// Above *pointer is an operator. It means, "Give me the value at this memory address".
}

/*
As we see above "*Type" is completely different from "*PointerVariable"
	1. *type is a type description.
	2. *pointer is an operator.
Also, the two operators:
	1. Turn address into value with "*address"
	2. Turn value into address with "&value"
*/
