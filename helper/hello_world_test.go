package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldFelix(t *testing.T) {
	result := HelloWorld("felixx")

	if result != "Hello world felix" {
		t.Error("Result must be Hello world felix")
	}

	fmt.Println("TestHelloWorldFelix done")
}

func TestHelloWorldXilef(t *testing.T) {
	result := HelloWorld("xileff")

	if result != "Hello world xilef" {
		t.Fatal("Hello world xilef")
	}

	fmt.Println("TestHelloWorldXilef done")
}
