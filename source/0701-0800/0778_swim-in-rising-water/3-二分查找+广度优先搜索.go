package main

import ()

func main() {

}

// leetcode778_水位上升的泳池中游泳
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func swimInWater(grid [][]int) int {
	n := len(grid)
	left, right := 0, n*n-1
	res := 0
	for left <= right {
		mid := left + (right-left)/2 // 二分枚举最大值
		if mid < grid[0][0] {
			left = mid + 1
			continue
		}
		queue := make([][2]int, 0)
		queue = append(queue, [2]int{0, 0})
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}
		for len(queue) > 0 {
			a, b := queue[0][0], queue[0][1]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				x, y := a+dx[i], b+dy[i]
				if 0 <= x && x < n && 0 <= y && y < n &&
					visited[x][y] == false && grid[x][y] <= mid {
					queue = append(queue, [2]int{x, y})
					visited[x][y] = true
				}
			}
		}
		if visited[n-1][n-1] == true { // 缩小范围
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
