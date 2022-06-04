package main

import "sort"

func main() {

}

// leetcode2250_统计包含每个点的矩形数目
func countRectangles(rectangles [][]int, points [][]int) []int {
	n := len(points)
	res := make([]int, n)
	arr := make([][]int, 101)
	for i := 0; i < len(rectangles); i++ {
		x, y := rectangles[i][0], rectangles[i][1]
		arr[y] = append(arr[y], x)
	}
	for i := 0; i < 101; i++ {
		sort.Ints(arr[i])
	}
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		for j := y; j < 101; j++ {
			total := len(arr[j]) - sort.SearchInts(arr[j], x) // 总和-不满足要求的点
			res[i] = res[i] + total                           // 累加
		}
	}
	return res
}
