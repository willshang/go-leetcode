package main

func main() {

}

// leetcode1391_检查网格中是否存在有效路径
func hasValidPath(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	arr := make([][]bool, 3*n)
	for i := 0; i < 3*n; i++ {
		arr[i] = make([]bool, 3*m)
	}
	for i := 0; i < n; i++ { // 放大处理
		for j := 0; j < m; j++ {
			x, y := 3*i, 3*j     // 9宫格的左上角坐标
			arr[x+1][y+1] = true // 换成9宫格后的中心点
			switch grid[i][j] {
			case 1:
				arr[x+1][y], arr[x+1][y+2] = true, true
			case 2:
				arr[x][y+1], arr[x+2][y+1] = true, true
			case 3:
				arr[x+1][y], arr[x+2][y+1] = true, true
			case 4:
				arr[x+1][y+2], arr[x+2][y+1] = true, true
			case 5:
				arr[x+1][y], arr[x][y+1] = true, true
			case 6:
				arr[x][y+1], arr[x+1][y+2] = true, true
			}
		}
	}
	return dfs(arr, 1, 1, 3*n-2, 3*m-2) // 从0,0的中心，走到n-1,m-1的中心
}

func dfs(arr [][]bool, i, j int, x, y int) bool {
	if i == x && j == y {
		return true
	}
	if i < 0 || i >= len(arr) || j < 0 || j >= len(arr[0]) || arr[i][j] == false {
		return false
	}
	arr[i][j] = false
	return dfs(arr, i-1, j, x, y) || dfs(arr, i+1, j, x, y) ||
		dfs(arr, i, j-1, x, y) || dfs(arr, i, j+1, x, y)
}
