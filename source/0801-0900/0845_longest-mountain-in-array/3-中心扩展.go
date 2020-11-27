package main

import "fmt"

func main() {
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}

func longestMountain(A []int) int {
	n := len(A)
	res := 0
	for i := 1; i < n-1; i++ {
		left, right := 0, 0
		for j := i - 1; j >= 0; j-- {
			if A[j] < A[j+1] {
				left++
			} else {
				break
			}
		}
		for j := i + 1; j < n; j++ {
			if A[j] < A[j-1] {
				right++
			} else {
				break
			}
		}
		if left > 0 && right > 0 {
			res = max(res, left+right+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
