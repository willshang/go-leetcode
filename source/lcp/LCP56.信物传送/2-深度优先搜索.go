package main

import "fmt"

func main() {
	fmt.Println(conveyorBelt([]string{">>v", "v^<", "<><"}, []int{0, 1}, []int{2, 0}))
}

// 顺时针：上右下左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var dir = []byte{'>', 'v', '<', '^'} // 注意对准坐标方向

var arr [][]int
var n, m int

func conveyorBelt(matrix []string, start []int, end []int) int {
	n, m = len(matrix), len(matrix[0])
	total := n * m
	arr = make([][]int, n) // 到达下标i,j的最少次数
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arr[i][j] = total
		}
	}
	dfs(matrix, start[0], start[1], 0)
	return arr[end[0]][end[1]]
}

func dfs(matrix []string, x, y int, count int) {
	if x < 0 || x >= n || y < 0 || y >= m {
		return
	}
	if count >= arr[x][y] {
		return
	}
	arr[x][y] = count
	index := 0
	for i := 0; i < 4; i++ {
		if matrix[x][y] == dir[i] {
			index = i
			break
		}
	}
	dfs(matrix, x+dx[index], y+dy[index], count) // 先走不变换方向：次数不变
	for i := 0; i < 4; i++ {
		newX, newY := x+dx[i], y+dy[i]
		if matrix[x][y] == dir[i] {
			continue
		}
		dfs(matrix, newX, newY, count+1) // 不同次数+1
	}
}
