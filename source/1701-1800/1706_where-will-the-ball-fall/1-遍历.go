package main

func main() {

}

// leetcode1706_球会落何处
func findBall(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	res := make([]int, m)
	for i := 0; i < m; i++ {
		res[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if res[j] != -1 {
				if res[j] != m-1 && grid[i][res[j]] == 1 && grid[i][res[j]+1] == 1 { // 向右侧 \
					res[j] = res[j] + 1 // 坐标向右+1
				} else if res[j] != 0 && grid[i][res[j]] == -1 && grid[i][res[j]-1] == -1 { // 向左侧 /
					res[j] = res[j] - 1 // 坐标向左-1
				} else {
					res[j] = -1
				}
			}
		}
	}
	return res
}
