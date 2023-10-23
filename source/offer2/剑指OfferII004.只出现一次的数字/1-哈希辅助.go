package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

// leetcode137_只出现一次的数字II
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
	return 0
}
