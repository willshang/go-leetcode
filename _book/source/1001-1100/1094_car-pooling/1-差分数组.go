package main

import "fmt"

func main() {
	arr := [][]int{
		{2, 1, 5},
		{3, 5, 7},
	}
	fmt.Println(carPooling(arr, 4))
}

// leetcode1094_拼车
func carPooling(trips [][]int, capacity int) bool {
	arr := make([]int, 1001)
	for i := 0; i < len(trips); i++ {
		start := trips[i][1]
		end := trips[i][2]
		count := trips[i][0]
		arr[start] = arr[start] + count
		arr[end] = arr[end] - count
	}
	total := 0
	for i := 0; i < len(arr); i++ {
		total = total + arr[i]
		if total > capacity {
			return false
		}
	}
	return true
}
