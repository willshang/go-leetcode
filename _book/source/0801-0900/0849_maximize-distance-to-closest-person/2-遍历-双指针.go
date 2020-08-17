package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 0, 0, 0, 1, 0, 1}
	fmt.Println(maxDistToClosest(arr))
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))
}

func maxDistToClosest(seats []int) int {
	n := len(seats)
	res := 0
	left := -1
	right := 0
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			left = i
		} else {
			// 找到右边有人的位置
			for (right < n && seats[right] == 0) || right < i {
				right++
			}
			leftLen := 0
			rightLen := 0
			if left == -1 {
				leftLen = n
			} else {
				leftLen = i - left
			}
			if right == n {
				rightLen = n
			} else {
				rightLen = right - i
			}
			if min(leftLen, rightLen) > res {
				res = min(leftLen, rightLen)
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
