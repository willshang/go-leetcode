package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

// leetcode1004_最大连续1的个数III
func longestOnes(A []int, K int) int {
	res := 0
	left, right := 0, 0
	count := 0
	for right = 0; right < len(A); right++ {
		if A[right] == 0 {
			count++
		}
		for count > K {
			if A[left] == 0 {
				count--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
