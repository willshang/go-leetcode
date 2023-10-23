package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

// 剑指OfferII119.最长连续序列
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]-1]; !ok {
			cur := nums[i]
			count := 1
			for m[cur+1] == true {
				count = count + 1
				cur = cur + 1
			}
			res = max(res, count)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
