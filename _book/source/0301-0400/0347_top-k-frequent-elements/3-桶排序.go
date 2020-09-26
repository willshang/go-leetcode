package main

import (
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}

// leetcode347_前K个高频元素
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	arr := make([][]int, len(nums)+1)
	temp := make(map[int][]int)
	for key, value := range m {
		temp[value] = append(temp[value], key)
		arr[value] = append(arr[value], key)
	}
	res := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		// 避免出现0=>x次的情况
		if _, ok := temp[i]; ok {
			for j := 0; j < len(arr[i]); j++ {
				k--
				if k < 0 {
					break
				}
				res = append(res, arr[i][j])
			}
		}
	}
	return res
}
