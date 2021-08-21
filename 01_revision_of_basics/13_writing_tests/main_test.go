package main

import "testing"

// Testing does requires a module: `go mod init github.com/kmrgaurav11235/13_writing_tests`
// Manual test
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have.") // Throw testing errors like this
	}
}

// Run `go test` or `go test -v` to run the tests
func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have.") // Throw testing errors like this
	}
}
