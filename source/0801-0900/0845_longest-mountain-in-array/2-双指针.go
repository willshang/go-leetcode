package main

import "fmt"

func main() {
	fmt.Println(longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
}

func longestMountain(A []int) int {
	n := len(A)
	left := 0
	res := 0
	for left+2 < n {
		// left指向左侧山脚, right寻找右侧山脚
		right := left + 1
		if A[left] < A[left+1] {
			for right+1 < n && A[right] < A[right+1] {
				right++
			}
			if right+1 < n && A[right] > A[right+1] {
				for right+1 < n && A[right] > A[right+1] {
					right++
				}
				if right-left+1 > res {
					res = right - left + 1
				}
			} else {
				right++
			}
		}
		left = right
	}
	return res
}
