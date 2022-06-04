package main

import (
	"sort"
)

func main() {

}

func countRectangles(rectangles [][]int, points [][]int) []int {
	n := len(points)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		points[i] = append(points[i], i) // 添加下标
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] > points[j][0] // 横坐标排序
	})
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i][0] > rectangles[j][0] // 横坐标排序
	})
	start := 0
	arr := make([]int, 101)
	for i := 0; i < n; i++ {
		x, y, index := points[i][0], points[i][1], points[i][2]
		for ; start < len(rectangles) && x <= rectangles[start][0]; start++ {
			arr[rectangles[start][1]]++ // 把纵坐标次数+1
		}
		for j := y; j < 101; j++ { // 遍历大于当前纵坐标
			res[index] = res[index] + arr[j] // 累加次数
		}
	}
	return res
}
