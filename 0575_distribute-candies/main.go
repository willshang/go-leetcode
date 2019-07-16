package main

import "fmt"

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(candies))
}
func distributeCandies(candies []int) int {
	n := len(candies)

	r := make(map[int]bool, n)
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
