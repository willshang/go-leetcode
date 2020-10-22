package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i < len(nums) && nums[i] != target {
		i++
	}
	for j >= 0 && nums[j] != target {
		j--
	}
	if i > j {
		return 0
	}
	return j - i + 1
}
