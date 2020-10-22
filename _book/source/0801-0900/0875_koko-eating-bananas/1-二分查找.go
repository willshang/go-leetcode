package main

import "fmt"

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

// leetcode875_爱吃香蕉的珂珂
func minEatingSpeed(piles []int, H int) int {
	maxValue := piles[0]
	for i := 1; i < len(piles); i++ {
		maxValue = max(maxValue, piles[i])
	}
	left, right := 1, maxValue
	for left < right {
		mid := left + (right-left)/2
		if judge(piles, mid, H) == true {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func judge(piles []int, speed int, H int) bool {
	total := 0
	for i := 0; i < len(piles); i++ {
		total = total + piles[i]/speed
		if piles[i]%speed > 0 {
			total = total + 1
		}
	}
	return total > H
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
