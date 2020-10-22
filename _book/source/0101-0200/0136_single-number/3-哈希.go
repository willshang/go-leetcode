package main

import "fmt"

func main() {
	arr := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}

// leetcode 136 只出现一次的数字
func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}
