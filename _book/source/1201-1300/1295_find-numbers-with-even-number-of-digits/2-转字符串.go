package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
}

// leetcode1295_统计位数为偶数的数字
func findNumbers(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		value := strconv.Itoa(nums[i])
		if len(value)%2 == 0 {
			res++
		}
	}
	return res
}
