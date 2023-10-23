package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrangeCoins(9))
}

// leetcode441_排列硬币
func arrangeCoins(n int) int {
	i := 1
	for i <= n {
		n = n - i
		i++
	}
	return i - 1
}
