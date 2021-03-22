package main

func main() {

}

func findBall(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	res := make([]int, m)
	for i := 0; i < m; i++ {
		res[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if res[j] == -1 {
				continue
			}
			if (grid[i][res[j]] == 1 && (res[j] == m-1 || grid[i][res[j]+1] == -1)) ||
				(grid[i][res[j]] == -1 && (res[j] == 0 || grid[i][res[j]-1] == 1)) {
				res[j] = -1
				continue
			}
			res[j] = res[j] + grid[i][res[j]]
		}
	}
	return res
}
