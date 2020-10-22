package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			res = append(res, abs(nums[i]))
		} else {
			nums[index] = -nums[index]
		}
	}
	return res
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
