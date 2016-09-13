package fizzlab_test

import (
	"testing"

	"github.com/girikuncoro/fizzlab"
)

var useCases = [...]struct {
	number      int
	expectation string
}{
	{0, ""},
	{1, ""},
	{3, "fizz"},
	{5, "buzz"},
	{9, "fizz"},
	{20, "buzz"},
}

func TestFizzbuzz(t *testing.T) {
	for _, useCase := range useCases {
		actual := fizzlab.Fizzbuzz(useCase.number)

		if actual != useCase.expectation {
			t.Fatalf("Expecting %s but got %s", useCase.expectation, actual)
		}
	}
}

func TestMain(t *testing.T) {
	var (
		expected = "12fizz4buzzfizz78fizzbuzz11fizz1314fizz1617fizz19buzz"
		actual   = fizzlab.Main()
	)

	if actual != expected {
		t.Fatalf("Expecting %s but got %s", expected, actual)
	}
}
