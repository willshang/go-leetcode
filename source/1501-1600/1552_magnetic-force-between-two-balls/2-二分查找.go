package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
}

// leetcode1552_两球之间的磁力
func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	maxValue := (position[n-1] - position[0]) / (m - 1) // 最大值
	minValue := 1                                       // 最小值
	left, right := minValue, maxValue
	res := 1
	for left <= right {
		mid := left + (right-left)/2
		// 满足条件，left=mid+1尝试最大
		if check(position, mid, m) {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}

func check(arr []int, value int, m int) bool {
	count := 1
	prev := 0
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[prev] >= value {
			count++
			prev = i
		}
	}
	return count >= m
}
