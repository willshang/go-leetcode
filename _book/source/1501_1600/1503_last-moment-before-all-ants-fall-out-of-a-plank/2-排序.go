package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getLastMoment(7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}))
}

// 2只蚂蚁相遇=>两只蚂蚁都不改变移动方向=>求离终点最远距离
func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	if len(left) == 0 {
		return n - right[0]
	}
	if len(right) == 0 {
		return left[len(left)-1]
	}
	if n-right[0] > left[len(left)-1] {
		return n - right[0]
	}
	return left[len(left)-1]
}
