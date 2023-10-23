package main

func main() {

}

// leetcode1631_最小体力消耗路径
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	left, right := 0, 1000000
	res := 0
	for left <= right {
		mid := left + (right-left)/2 // 二分枚举最大值
		queue := make([][2]int, 0)
		queue = append(queue, [2]int{0, 0})
		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, m)
		}
		for len(queue) > 0 {
			a, b := queue[0][0], queue[0][1]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				x, y := a+dx[i], b+dy[i]
				if 0 <= x && x < n && 0 <= y && y < m &&
					visited[x][y] == false && abs(heights[a][b]-heights[x][y]) <= mid {
					queue = append(queue, [2]int{x, y})
					visited[x][y] = true
				}
			}
		}
		if visited[n-1][m-1] == true { // 缩小范围
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
