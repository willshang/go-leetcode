package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}

func findKthPositive(arr []int, k int) int {
	start := 1
	i := 0
	for {
		if i < len(arr) && arr[i] == start {
			i++
		} else {
			k--
			if k == 0 {
				return start
			}
		}
		start++
	}
}
