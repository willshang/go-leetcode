package main

import "fmt"

func main() {
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(1))
}

func mySqrt(x int) int {
	result := 0
	for i := 1; i <= x/i; i++ {
		if i*i == x {
			return i
		}
		result = i
	}
	return result
}
