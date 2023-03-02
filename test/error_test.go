package test

import (
	"fmt"
	"testing"
	index "unitTest"
)

func TestHelloWorldErrorFail(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		t.Fail()
	}

	fmt.Println("Finished Executd Fail Test")
}

func TestHelloWorldErrorFailNow(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		t.FailNow()
	}

	fmt.Println("Finished Executd Fail Now Test")
}

func TestHelloWorldErrors(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		t.Error("Should be Reihan")
	}

	fmt.Println("Finished Executd Error Test")
}

func TestHelloWorldErrorFatal(t *testing.T) {
	result := index.HelloWorld("Reihan")

	if result != "Hello Reihan" {
		t.Fatal("Should be Reihan")
	}

	fmt.Println("Finished Executd Fatal Test")
}
