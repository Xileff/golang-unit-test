package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("felix")

	if result != "Hello world felix" {
		panic("Result is not Hello world felix")
	}
}
