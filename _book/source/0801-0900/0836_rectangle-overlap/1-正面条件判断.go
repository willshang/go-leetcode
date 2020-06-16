package main

import "fmt"

func main() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 满足条件
	if rec1[1] < rec2[3] && rec1[0] < rec2[2] &&
		rec2[1] < rec1[3] && rec2[0] < rec1[2] {
		return true
	}
	return false
}
