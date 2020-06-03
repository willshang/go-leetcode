package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	result := 0
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
