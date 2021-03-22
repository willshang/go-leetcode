package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}

// leetcode1512_好数对的数目
func numIdenticalPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return res
}
