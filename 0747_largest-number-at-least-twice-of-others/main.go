package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
	fmt.Println(dominantIndex([]int{1,0}))
}
func dominantIndex(nums []int) int {
	n := len(nums)
	if n == 1{
		return 0
	}

	first, second := 0,1
	if nums[first] < nums[second]{
		first,second = second,first
	}

	for i := 2; i < n; i++{
		if nums[first] < nums[i]{
			first,second = i,first
		}else if nums[second] < nums[i]{
			second = i
		}
	}

	if nums[second] == 0 || nums[first] / nums[second] >= 2{
		return first
	}
	return -1
}