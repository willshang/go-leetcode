package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(10))
}

// 剑指OfferII072.求平方根
func mySqrt(x int) int {
	result := x
	for result*result > x {
		result = (result + x/result) / 2
	}
	return result
}
