package main

import (
	"fmt"

	"github.com/kmrgaurav11235/mypackage/helpers"
)

// go mod init github.com/kmrgaurav11235/mypackage to generate the go.mod file

func main() {
	var myVar helpers.UserType
	fmt.Println(myVar.UserName)
}
