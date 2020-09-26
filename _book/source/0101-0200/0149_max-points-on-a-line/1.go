package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}
	res := 2
	for i := 0; i < n; i++ {
		// k = (y2 - y1) / (x2 - x1)
		// b = y1 - kx1

	}
}
