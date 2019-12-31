package main

import "fmt"

func main() {
	arr := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}

// leetcode 136 只出现一次的数字
func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res ^ n
	}
	return res
}
