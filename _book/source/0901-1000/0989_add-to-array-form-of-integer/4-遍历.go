package main

import "fmt"

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
}

// leetcode989_数组形式的整数加法
func addToArrayForm(A []int, K int) []int {
	i := len(A) - 1
	res := make([]int, 0)
	for i >= 0 || K > 0 {
		if i >= 0 {
			K = K + A[i]
		}
		res = append([]int{K % 10}, res...)
		K = K / 10
		i--
	}
	return res
}
