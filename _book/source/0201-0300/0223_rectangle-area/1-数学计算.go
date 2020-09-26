package main

import "fmt"

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}

// leetcode223_矩形面积
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	left, right := max(A, E), min(C, G)
	bottom, top := max(B, F), min(D, H)
	area1, area2 := (C-A)*(D-B), (G-E)*(H-F)
	if left < right && bottom < top {
		return area1 + area2 - (right-left)*(top-bottom)
	}
	return area1 + area2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
