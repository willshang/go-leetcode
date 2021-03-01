package main

func main() {

}

// leetcode1267_统计参与通信的服务器
func countServers(grid [][]int) int {
	a := make(map[int]int) // 行
	b := make(map[int]int) // 列
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				a[i]++
				b[j]++
			}
		}
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && (a[i] > 1 || b[j] > 1) {
				res++
			}
		}
	}
	return res
}
