package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2}))
}

// 程序员面试金典17.10_主要元素
func majorityElement(nums []int) int {
	m := make(map[int]int)
	result := -1
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
		if m[v] > (len(nums) / 2) {
			result = v
		}
	}
	return result
}
