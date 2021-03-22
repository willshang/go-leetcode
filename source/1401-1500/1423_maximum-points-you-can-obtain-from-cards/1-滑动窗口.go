package main

import "fmt"

func main() {
	//fmt.Println(maxScore([]int{2, 2, 2}, 2))
	//fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}

// leetcode1423_可获得的最大点数
func maxScore(cardPoints []int, k int) int {
	res := 0
	left := 0
	for i := 0; i < k; i++ {
		left = left + cardPoints[i]
	}
	res = left
	right := 0
	count := k
	for i := len(cardPoints) - 1; i >= len(cardPoints)-k; i-- {
		right = right + cardPoints[i]
		left = left - cardPoints[count-1]
		res = max(res, left+right)
		count--
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
