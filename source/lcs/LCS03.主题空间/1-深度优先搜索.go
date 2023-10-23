package main

func main() {

}

// leetcode_lcs03_主题空间
var count int
var flag bool

func largestArea(grid []string) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if '1' <= grid[i][j] && grid[i][j] <= '5' {
				count, flag = 0, true
				dfs(grid, i, j, grid[i][j])
				if flag == true && count > res {
					res = count
				}
			}
		}
	}
	return res
}

func dfs(grid []string, i, j int, char byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' { // 不满足条件
		flag = false
		return
	}
	if grid[i][j] != char { // 不一致
		return
	}
	count++
	arr := []byte(grid[i])
	arr[j] = arr[j] + 5 // 变换为其它的数
	grid[i] = string(arr)
	dfs(grid, i+1, j, char)
	dfs(grid, i-1, j, char)
	dfs(grid, i, j+1, char)
	dfs(grid, i, j-1, char)
	return
}
