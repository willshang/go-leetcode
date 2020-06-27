package main

import (
	"fmt"
)

func main() {
	fmt.Println(distributeCandies(7, 4))
}

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	i := 0
	count := 0
	for candies > 0 {
		count++
		if candies >= count {
			res[i%num_people] += count
		} else {
			res[i%num_people] += candies
		}
		i++
		candies = candies - count
	}
	return res
}
