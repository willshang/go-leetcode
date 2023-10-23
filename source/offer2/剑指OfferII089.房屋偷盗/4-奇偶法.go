package main

import "fmt"

func main() {
	var nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
