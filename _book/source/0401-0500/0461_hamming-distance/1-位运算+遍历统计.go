package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(4, 1))
}

func hammingDistance(x int, y int) int {
	x = x ^ y
	res := 0
	for x > 0 {
		if x&1 == 1 {
			res++
		}
		x = x >> 1
	}
	return res
}
