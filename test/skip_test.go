package test

import (
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
	index "unitTest"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di Mac")
	}

	result := index.HelloWorld("Reihan")
	require.Equal(t, "Hello Reihan", result)
}
