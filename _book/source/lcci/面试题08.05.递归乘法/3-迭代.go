package main

import "fmt"

func main() {
	fmt.Println(multiply(3, 4))
}

func multiply(A int, B int) int {
	res := 0
	for B != 0 {
		if B%2 == 1 {
			res = res + A
		}
		A = A + A
		B = B >> 1
	}
	return res
}
