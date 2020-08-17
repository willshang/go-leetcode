package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
}

// 剑指offer_面试题56-I_数组中数字出现的次数
func singleNumbers(nums []int) []int {
	res := make([]int, 2)
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = temp ^ nums[i]
	}
	last := 1
	for temp&last == 0 {
		last = last << 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]&last == 0 {
			res[0] = res[0] ^ nums[i]
		} else {
			res[1] = res[1] ^ nums[i]
		}
	}
	return res
}
