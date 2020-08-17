package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

// leetcode1431_拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	max := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}
	return res
}
