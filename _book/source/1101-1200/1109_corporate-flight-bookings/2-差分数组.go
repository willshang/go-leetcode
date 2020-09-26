package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	}
	fmt.Println(corpFlightBookings(arr, 5))
}

// leetcode1109_航班预订统计
func corpFlightBookings(bookings [][]int, n int) []int {
	arr := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		start := bookings[i][0]
		end := bookings[i][1]
		count := bookings[i][2]
		arr[start-1] = arr[start-1] + count
		if end < n {
			arr[end] = arr[end] - count
		}
	}
	for i := 1; i < n; i++ {
		arr[i] = arr[i] + arr[i-1]
	}
	return arr
}
