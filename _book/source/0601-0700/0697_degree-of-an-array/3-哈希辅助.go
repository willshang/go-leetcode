package main

import "fmt"

func main() {
	//fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
}

// leetcode697_数组的度
func findShortestSubArray(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	res := 0
	maxLen := 0
	m := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}
	for _, v := range m {
		if len(v) > maxLen {
			maxLen = len(v)
			res = v[len(v)-1] - v[0] + 1
		} else if len(v) == maxLen && v[len(v)-1]-v[0]+1 < res {
			res = v[len(v)-1] - v[0] + 1
		}
	}
	return res
}
