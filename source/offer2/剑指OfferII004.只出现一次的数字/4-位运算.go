package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, -3, 2}))
}

func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = (a ^ nums[i]) & (^b) // a：保留出现1次的数
		b = (b ^ nums[i]) & (^a) // b：保留出现2次的数
	}
	return a // 最后返回只出现1次的数
}
