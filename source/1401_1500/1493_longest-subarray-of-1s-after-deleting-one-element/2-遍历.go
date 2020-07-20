package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
}

// leetcode1493_删掉一个元素以后全为 1 的最长子数组
func longestSubarray(nums []int) int {
	res := 0
	p, q := 0, 0 // q=>中间有一个“非1”的和, p=>连续1的和
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			q = p
			p = 0
		} else {
			p++
			q++
		}
		if q > res {
			res = q
		}
	}
	if res == len(nums) {
		return res - 1
	}
	return res
}
