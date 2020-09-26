package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

// leetcode229_求众数II
func majorityElement(nums []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}
