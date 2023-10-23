package main

import (
	"fmt"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}

func heightChecker(heights []int) int {
	arr := make([]int, 101)
	for i := 0; i < len(heights); i++ {
		arr[heights[i]]++
	}
	res := 0
	j := 0
	for i := 1; i <= 100; i++ {
		for arr[i] > 0 {
			if heights[j] != i {
				res++
			}
			arr[i]--
			j++
		}
	}
	return res
}
