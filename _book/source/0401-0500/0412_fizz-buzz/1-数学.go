package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

// leetcode412_FizzBuzz
func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		x := i + 1
		switch {
		case x%15 == 0:
			ret[i] = "FizzBuzz"
		case x%5 == 0:
			ret[i] = "Buzz"
		case x%3 == 0:
			ret[i] = "Fizz"
		default:
			ret[i] = strconv.Itoa(x)
		}
	}
	return ret
}
