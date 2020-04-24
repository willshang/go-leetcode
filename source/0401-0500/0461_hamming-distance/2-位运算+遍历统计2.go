package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(4, 1))
}

// leetcode461_汉明距离
func hammingDistance(x int, y int) int {
	x = x ^ y
	res := 0
	for x > 0 {
		res++
		x = x & (x - 1)
	}
	return res
}
