package main

import "fmt"

func main() {
	fmt.Println(conveyorBelt([]string{">>v", "v^<", "<><"}, []int{0, 1}, []int{2, 0}))
}

// leetcode_lcp56_信物传送
// 顺时针：上右下左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var dir = []byte{'>', 'v', '<', '^'} // 注意对准坐标方向

func conveyorBelt(matrix []string, start []int, end []int) int {
	n, m := len(matrix), len(matrix[0])
	total := n * m
	arr := make([][]int, n) // 到达下标i,j的最少次数
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			arr[i][j] = total
		}
	}
	arr[start[0]][start[1]] = 0
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{start[0], start[1]})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		x, y := node[0], node[1]
		for i := 0; i < 4; i++ {
			newX, newY := x+dx[i], y+dy[i]
			if 0 <= newX && newX < n && 0 <= newY && newY < m {
				if matrix[x][y] == dir[i] { // 不变换方向
					if arr[newX][newY] > arr[x][y] {
						arr[newX][newY] = arr[x][y] // 更新次数
						queue = append(queue, [2]int{newX, newY})
					}
				} else { // 变换方向，次数+1
					if arr[newX][newY] > arr[x][y]+1 {
						arr[newX][newY] = arr[x][y] + 1 // 更新次数
						queue = append(queue, [2]int{newX, newY})
					}
				}
			}
		}
	}
	return arr[end[0]][end[1]]
}
