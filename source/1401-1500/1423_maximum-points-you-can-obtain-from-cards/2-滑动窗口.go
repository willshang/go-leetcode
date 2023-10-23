package main

import "fmt"

func main() {
	//fmt.Println(maxScore([]int{2, 2, 2}, 2))
	//fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}

func maxScore(cardPoints []int, k int) int {
	res := 0
	n := len(cardPoints)
	window := 0
	sum := 0
	for i := 0; i < n-k; i++ {
		sum = sum + cardPoints[i]
		window = window + cardPoints[i]
	}
	res = window
	count := 0
	for i := n - k; i < n; i++ {
		sum = sum + cardPoints[i]
		window = window + cardPoints[i] - cardPoints[count]
		res = min(res, window)
		count++
	}
	return sum - res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
