package test

import (
	"testing"
	index "unitTest"
)

func TestHelloWorld(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		//	error
		panic("Result is not Hello 'Reihan'")
	}
}
