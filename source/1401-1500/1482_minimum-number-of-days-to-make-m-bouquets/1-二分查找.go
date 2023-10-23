package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
}

// leetcode1482_制作m束花所需的最少天数
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	minValue, maxValue := bloomDay[0], bloomDay[0]
	for i := 1; i < len(bloomDay); i++ {
		if bloomDay[i] > maxValue {
			maxValue = bloomDay[i]
		}
		if bloomDay[i] < minValue {
			minValue = bloomDay[i]
		}
	}
	left, right := minValue, maxValue
	for left < right {
		mid := left + (right-left)/2
		count := judge(bloomDay, mid, k)
		if count >= m {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(bloomDay []int, mid int, k int) int {
	total := 0
	count := 0
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] <= mid {
			total++
		} else {
			total = 0
		}
		if total == k {
			count++
			total = 0
		}
	}
	return count
}
