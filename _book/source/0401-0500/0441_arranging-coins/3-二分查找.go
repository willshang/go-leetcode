package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrangeCoins(9))
}

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if mid*(mid+1)/2 < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left*(left+1)/2 == n {
		return left
	}
	return left - 1
}
