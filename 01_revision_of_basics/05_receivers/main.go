package main

import "log"

type myStruct struct {
	FirstName string
}

func (m myStruct) getFirstName() string {
	// (m myStruct) is a Receiver. Any variable of type "myStruct" now gets access to this "getFirstName" method
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Pierre"

	myVar2 := myStruct{
		FirstName: "Conseil",
	}

	log.Println("myVar is set to", myVar.getFirstName())
	log.Println("myVar2 is set to", myVar2.getFirstName())
}
