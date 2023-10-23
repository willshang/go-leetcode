package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

// leetcode260_只出现一次的数字III
func singleNumber(nums []int) []int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a = a ^ nums[i]
	}
	b := a & (-a) // x & (-x) 是保留位中最右边 1 ，且将其余的 1 设位 0 的方法。
	value := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&b == 0 {
			value = value ^ nums[i]
		}
	}
	return []int{value, a ^ value}
}
