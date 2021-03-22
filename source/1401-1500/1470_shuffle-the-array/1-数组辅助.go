package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}

// leetcode1470_重新排列数组
func shuffle(nums []int, n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, nums[i], nums[i+n])
	}
	return res
}
