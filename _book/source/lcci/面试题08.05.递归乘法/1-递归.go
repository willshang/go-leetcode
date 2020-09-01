package main

import "fmt"

func main() {
	fmt.Println(multiply(3, 4))
}

func multiply(A int, B int) int {
	if B == 0 {
		return 0
	}
	return multiply(A, B-1) + A
}
