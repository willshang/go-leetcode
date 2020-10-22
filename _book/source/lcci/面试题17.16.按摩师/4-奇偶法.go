package main

import "fmt"

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
}

func massage(nums []int) int {
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
