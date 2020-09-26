package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	res := 0
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] > 0 {
			continue
		}
		left := m[nums[i]-1]
		right := m[nums[i]+1]
		sum := left + 1 + right
		res = max(res, sum)
		m[nums[i]] = sum
		m[nums[i]-left] = sum
		m[nums[i]+right] = sum
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
