package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before unit tests")

	m.Run()

	fmt.Println("After unit tests")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("felix")

	require.Equal(t, "Hello world felix", result, "Result must be Hello world felix")

	fmt.Println("TestHelloWorldRequire done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("felix")

	assert.Equal(t, "Hello world felix", result, "Result must be Hello world felix")

	fmt.Println("TestHelloWorldAssert done")
}
