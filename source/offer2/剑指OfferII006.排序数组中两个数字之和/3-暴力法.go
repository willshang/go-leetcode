package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
