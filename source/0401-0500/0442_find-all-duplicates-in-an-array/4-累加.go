package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		index := nums[i]%(n+1) - 1
		nums[index] = nums[index] + (n + 1)
	}
	for i := 0; i < n; i++ {
		if nums[i]/(n+1) == 2 {
			res = append(res, i+1)
		}
	}
	return res
}
