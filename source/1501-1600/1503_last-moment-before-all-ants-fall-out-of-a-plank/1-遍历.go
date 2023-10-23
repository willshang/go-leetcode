package main

import "fmt"

func main() {
	fmt.Println(getLastMoment(7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}))
}

// leetcode1503_所有蚂蚁掉下来前的最后一刻
// 2只蚂蚁相遇=>两只蚂蚁都不改变移动方向=>求离终点最远距离
func getLastMoment(n int, left []int, right []int) int {
	max := 0
	for i := 0; i < len(left); i++ {
		if left[i] > max {
			max = left[i]
		}
	}
	for i := 0; i < len(right); i++ {
		if n-right[i] > max {
			max = n - right[i]
		}
	}
	return max
}
