package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

func longestOnes(A []int, K int) int {
	res := 0
	left, right := 0, 0
	for right < len(A) {
		if A[right] == 1 {
			right++
		} else {
			if K > 0 {
				right++
				K--
			} else {
				res = max(res, right-left)
				if A[left] == 0 {
					K++
				}
				left++
			}
		}
	}
	res = max(res, right-left)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
