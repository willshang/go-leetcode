package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
}

// leetcode1184_公交站间的距离
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	x := 0
	y := 0
	for i := start; i != destination; i = (i + 1) % len(distance) {
		x = x + distance[i]
	}
	for i := destination; i != start; i = (i + 1) % len(distance) {
		y = y + distance[i]
	}
	if x > y {
		return y
	}
	return x
}
