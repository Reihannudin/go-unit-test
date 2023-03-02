package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	index "unitTest"
)

func TestHelloWorldAssertion(t *testing.T) {
	result := index.HelloWorld("Reihan")
	assert.Equal(t, "Hello Reihan", result)
	fmt.Println("Finished")
}

func TestHelloWorldRequired(t *testing.T) {
	result := index.HelloWorld("Reihan")
	require.Equal(t, "Hello Reihan", result)
	fmt.Println("Finished")
}
