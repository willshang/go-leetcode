package main

import "fmt"

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(candies))
}

// leetcode575_分糖果
func distributeCandies(candies []int) int {
	n := len(candies)
	r := make(map[int]bool)
	for _, c := range candies {
		r[c] = true
	}
	return min(len(r), n/2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
