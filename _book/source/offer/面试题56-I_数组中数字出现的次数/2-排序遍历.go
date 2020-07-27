package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(singleNumbers([]int{1, 2, 10, 4, 1, 4, 3, 3}))
}

func singleNumbers(nums []int) []int {
	res := make([]int, 0)
	sort.Ints(nums)
	count := 0
	for i := 0; i < len(nums)-2; {
		if nums[i] == nums[i+1] {
			i = i + 2
		}
		if i == len(nums)-1 {
			res = append(res, nums[i])
			return res
		}
		if nums[i] != nums[i+1] {
			res = append(res, nums[i])
			i++
			count++
			if count == 2 {
				return res
			}
		}
	}
	return res
}
