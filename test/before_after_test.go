package test

import (
	"fmt"
	"testing"
	index "unitTest"
)

func TestHelloBefore(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		//	error
		panic("Result is not Hello 'Reihan'")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before unit test")

	m.Run()

	fmt.Println("After unit test")
}

func TestHelloAfter(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		//	error
		panic("Result is not Hello 'Reihan'")
	}
}
