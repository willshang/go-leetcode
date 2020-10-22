package main

import "fmt"

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}

// leetcode1011_在D天内送达包裹的能力
func shipWithinDays(weights []int, D int) int {
	sum := weights[0]
	maxValue := weights[0]
	for i := 1; i < len(weights); i++ {
		sum = sum + weights[i]
		if weights[i] > maxValue {
			maxValue = weights[i]
		}
	}
	left, right := maxValue, sum
	for left < right {
		mid := left + (right-left)/2
		count := judge(weights, mid)
		if count <= D {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(weights []int, target int) int {
	total := 0
	count := 1
	for i := 0; i < len(weights); i++ {
		if total+weights[i] <= target {
			total = total + weights[i]
		} else {
			count++
			total = weights[i]
		}
	}
	return count
}
