package main

import (
	"fmt"
	"sort"
)

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(candies))
}

func distributeCandies(candies []int) int {
	length := len(candies)
	half := length / 2
	count := 1
	sort.Ints(candies)
	for i := 1; i < length; i++ {
		if candies[i] != candies[i-1] {
			count++
		}
	}
	if count >= half {
		return half
	}
	return count
}
