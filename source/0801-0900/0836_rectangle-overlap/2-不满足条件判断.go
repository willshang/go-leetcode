package main

import "fmt"

func main() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 不满足条件, rec2固定，rec1在rec2的方位
	// 左侧：rec1[2] <= rec2[0]
	// 右侧：rec1[0] >= rec2[2]
	// 上方：rec1[1] >= rec2[3]
	// 下方：rec1[3] <= rec2[1]
	if rec1[2] <= rec2[0] ||
		rec1[3] <= rec2[1] ||
		rec1[0] >= rec2[2] ||
		rec1[1] >= rec2[3] {
		return false
	}
	return true
}
