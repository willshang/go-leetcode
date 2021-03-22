package main

import "fmt"

func main() {
	fmt.Println(countOdds(3, 7))
}

// leetcode1523_在区间范围内统计奇数数目
func countOdds(low int, high int) int {
	if low%2 == 0 && high%2 == 0 {
		return (high - low) / 2
	}
	return (high-low)/2 + 1
}
