package fizzbuzz

import (
	"fmt"
)

func Fizzbuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "fizzbuzz"
	}

	if n%3 == 0 {
		return "fizz"
	}
	if n%5 == 0 {
		return "buzz"
	}
	return fmt.Sprintf("%v", n)
}
