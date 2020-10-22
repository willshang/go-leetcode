package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	x := 0
	sum := 0
	for i := 0; i < len(distance); i++ {
		sum = sum + distance[i]
		if start < destination {
			if i >= start && i < destination {
				x = x + distance[i]
			}
		} else {
			if i >= destination && i < start {
				x = x + distance[i]
			}
		}
	}
	if sum-x > x {
		return x
	}
	return sum - x
}
