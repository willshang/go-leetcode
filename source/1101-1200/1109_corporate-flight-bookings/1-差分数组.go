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

func corpFlightBookings(bookings [][]int, n int) []int {
	arr := make([]int, n+1)
	for i := 0; i < len(bookings); i++ {
		start := bookings[i][0] - 1
		end := bookings[i][1] - 1
		count := bookings[i][2]
		arr[start] = arr[start] + count
		arr[end+1] = arr[end+1] - count
	}
	res := make([]int, 0)
	total := 0
	for i := 0; i < n; i++ {
		total = total + arr[i]
		res = append(res, total)
	}
	return res
}
