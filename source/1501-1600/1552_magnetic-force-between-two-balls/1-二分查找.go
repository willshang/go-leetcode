package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxDistance([]int{1, 2, 3, 4, 7}, 3))
}

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	maxValue := position[n-1] - position[0] // 最大值
	minValue := position[n-1]               // 求最小值
	for i := 1; i < n; i++ {
		if minValue > position[i]-position[i-1] {
			minValue = position[i] - position[i-1]
		}
	}
	if m == 2 {
		return maxValue
	}
	left, right := minValue, maxValue
	for left <= right {
		mid := left + (right-left)/2
		if search(position, mid, m) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left - 1
}

func search(arr []int, value int, m int) bool {
	count := 1
	target := arr[0] + value
	for i := 0; i < len(arr); i++ {
		if arr[i] >= target {
			count++
			target = arr[i] + value
		}
	}
	return count >= m
}
