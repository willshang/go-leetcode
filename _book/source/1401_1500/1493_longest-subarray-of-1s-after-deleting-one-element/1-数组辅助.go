package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] == 1 {
			pre[i] = pre[i-1] + 1
		} else {
			pre[i] = 0
		}
	}
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] == 1 {
			suf[i] = suf[i+1] + 1
		} else {
			suf[i] = 0
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		var p, s int
		if i == 0 {
			p = 0
		} else {
			p = pre[i-1]
		}
		if i == n-1 {
			s = 0
		} else {
			s = suf[i+1]
		}
		if p+s > res {
			res = p + s
		}
	}
	return res
}
