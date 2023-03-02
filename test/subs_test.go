package test

import (
	"github.com/stretchr/testify/require"
	"testing"
	index "unitTest"
)

func TestSubTest(t *testing.T) {
	t.Run("Reihan", func(t *testing.T) {
		result := index.HelloWorld("Reihan")
		require.Equal(t, "Hello Reihan", result)
	})
	t.Run("Andrian", func(t *testing.T) {
		result := index.HelloWorld("Andrian")
		require.Equal(t, "Hello Andrian", result)
	})
}
