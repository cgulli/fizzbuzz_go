package main

import (
	"fmt"
	"exercises/fizzbuzz"
)

func main(){
	for n := 1; n <= 100; n++ {
		fmt.Println(fizzbuzz.Fizzbuzz(n))
	}
}
