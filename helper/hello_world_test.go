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

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Felix)",
			request:  "Felix",
			expected: "Hello world Felix",
		},
		{
			name:     "HelloWorld(Xilef)",
			request:  "Xilef",
			expected: "Hello world Xilef",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	fmt.Println("Start TestSubTest")

	t.Run("Felix", func(t *testing.T) {
		result := HelloWorld("felix")
		require.Equal(t, "Hello world felix", result, "Result must be Hello world felix")
	})

	t.Run("Xilef", func(t *testing.T) {
		result := HelloWorld("xilef")
		require.Equal(t, "Hello world xilef", result, "Result must be Hello world xilef")
	})

	fmt.Println("End TestSubTest")
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
