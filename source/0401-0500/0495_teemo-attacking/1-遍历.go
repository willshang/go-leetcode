package main

import "fmt"

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
}

// leetcode495_提莫攻击
func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	if len(timeSeries) == 0 {
		return 0
	}
	for i := 0; i < len(timeSeries)-1; i++ {
		res = res + min(timeSeries[i+1]-timeSeries[i], duration)
	}
	return res + duration
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
