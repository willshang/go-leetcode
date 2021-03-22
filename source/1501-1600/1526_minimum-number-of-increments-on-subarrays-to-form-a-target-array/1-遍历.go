package main

import "fmt"

func main() {
	fmt.Println(minNumberOperations([]int{1, 2, 3, 2, 1}))
}

// leetcode1526_形成目标数组的子数组最少增加次数
func minNumberOperations(target []int) int {
	res := target[0]
	for i := 1; i < len(target); i++ {
		res = res + max(target[i]-target[i-1], 0)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
