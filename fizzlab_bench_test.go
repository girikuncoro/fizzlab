package fizzlab_test

import (
	"testing"

	"github.com/girikuncoro/fizzlab"
)

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzlab.Fizzbuzz(i)
	}
}
