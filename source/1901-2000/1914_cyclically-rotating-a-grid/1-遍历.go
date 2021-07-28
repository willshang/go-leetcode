package main

func main() {

}

// leetcode1914_循环轮转矩阵
func rotateGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	count := min(n/2, m/2)
	for level := 0; level < count; level++ {
		arr := make([][3]int, 0)
		for i := level; i < n-1-level; i++ { // 左上=>左下
			arr = append(arr, [3]int{i, level, grid[i][level]})
		}
		for j := level; j < m-1-level; j++ { // 左下=>右下
			arr = append(arr, [3]int{n - 1 - level, j, grid[n-1-level][j]})
		}
		for i := n - 1 - level; i > level; i-- { // 右下=>右上
			arr = append(arr, [3]int{i, m - 1 - level, grid[i][m-1-level]})
		}
		for j := m - 1 - level; j > level; j-- { // 右上=>左上
			arr = append(arr, [3]int{level, j, grid[level][j]})
		}
		total := len(arr)
		step := k % total
		for i := 0; i < total; i++ {
			index := (i + total - step) % total
			a, b, c := arr[i][0], arr[i][1], arr[index][2]
			grid[a][b] = c
		}
	}
	return grid
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
