package main

func main() {

}

// leetcode934_最短的桥
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var queue [][2]int
var n int

func shortestBridge(grid [][]int) int {
	n = len(grid)
	queue = make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j)  // 深度优先搜索找边界
				return bfs(grid) // 广度优先搜索找距离
			}
		}
	}
	return 0
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || i >= n || j < 0 || j >= n {
		return
	}
	if grid[i][j] == 0 {
		queue = append(queue, [2]int{i, j}) // 加入边界
	} else if grid[i][j] == 1 {
		grid[i][j] = 2 // 标记为2
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			dfs(grid, x, y)
		}
	}
}

func bfs(grid [][]int) int {
	res := 0
	for len(queue) > 0 {
		res++
		length := len(queue)
		for i := 0; i < length; i++ {
			a, b := queue[i][0], queue[i][1]
			for j := 0; j < 4; j++ {
				x, y := a+dx[j], b+dy[j]
				if 0 <= x && x < n && 0 <= y && y < n {
					if grid[x][y] == 2 {
						continue
					}
					if grid[x][y] == 1 {
						return res
					}
					queue = append(queue, [2]int{x, y})
					grid[x][y] = 2
				}
			}
		}
		queue = queue[length:]
	}
	return res
}
