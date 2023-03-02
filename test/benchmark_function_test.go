package test

import (
	"testing"
	index "unitTest"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		index.HelloWorld("Eko")
	}
}
func BenchmarkHelloWorldReihannudin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		index.HelloWorld("Reihan")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Andrian", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			index.HelloWorld("Andrian")
		}
	})
	b.Run("Reihan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			index.HelloWorld("Reihan")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
		{
			name:    "EkoKurniawanKhannedy",
			request: "Eko Kurniawan Khannedy",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				index.HelloWorld(benchmark.request)
			}
		})
	}
}
