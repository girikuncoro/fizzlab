package fizzlab

import "fmt"

// Fizzbuzz blah
func Fizzbuzz(i int) string {
	if i == 0 {
		return ""
	} else if i%3 == 0 {
		return "fizz"
	} else if i%5 == 0 {
		return "buzz"
	}
	return ""
}

// Main function of fizzlab
func Main() string {
	res := ""
	for i := 1; i < 21; i++ {
		s := Fizzbuzz(i)
		if s != "" {
			res += s
		} else {
			res += fmt.Sprintf("%d", i)
		}
	}
	return res
}
