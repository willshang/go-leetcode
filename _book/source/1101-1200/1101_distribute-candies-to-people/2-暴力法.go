package main

import "fmt"

func main() {
	fmt.Println(distributeCandies(7, 4))
}

// leetcode1101_分糖果 II
func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	count := 1
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			if candies >= count {
				res[i] = res[i] + count
				candies = candies - count
			} else {
				res[i] = res[i] + candies
				candies = 0
			}
			count++
		}
	}
	return res
}
