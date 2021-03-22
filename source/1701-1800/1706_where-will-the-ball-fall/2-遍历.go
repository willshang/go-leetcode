package main

func main() {

}

func findBall(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	res := make([]int, 0)
	for j := 0; j < m; j++ { // 每列模拟
		index := j
		for i := 0; i < n; i++ {
			if (grid[i][index] == 1 && (index == m-1 || grid[i][index+1] == -1)) ||
				(grid[i][index] == -1 && (index == 0 || grid[i][index-1] == 1)) {
				index = -1
				break
			}
			index = index + grid[i][index]
		}
		res = append(res, index)
	}
	return res
}
