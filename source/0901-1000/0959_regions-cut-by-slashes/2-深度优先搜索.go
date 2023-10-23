package main

func main() {

}

// leetcode959_由斜杠划分区域
func regionsBySlashes(grid []string) int {
	n := len(grid)
	arr := make([][]int, 3*n)
	for i := 0; i < 3*n; i++ {
		arr[i] = make([]int, 3*n)
	}
	for i := 0; i < n; i++ { // 放大处理
		for j := 0; j < n; j++ {
			if grid[i][j] == '/' { // 9宫格
				arr[i*3+2][j*3] = 1
				arr[i*3+1][j*3+1] = 1
				arr[i*3][j*3+2] = 1
			} else if grid[i][j] == '\\' {
				arr[i*3+2][j*3+2] = 1
				arr[i*3+1][j*3+1] = 1
				arr[i*3][j*3] = 1
			}
		}
	}
	res := 0
	for i := 0; i < 3*n; i++ {
		for j := 0; j < 3*n; j++ {
			if arr[i][j] == 0 {
				dfs(arr, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(arr [][]int, i, j int) {
	if 0 <= i && i < len(arr) && 0 <= j && j < len(arr[0]) && arr[i][j] == 0 {
		arr[i][j] = 1
		dfs(arr, i+1, j)
		dfs(arr, i-1, j)
		dfs(arr, i, j-1)
		dfs(arr, i, j+1)
	}
}
