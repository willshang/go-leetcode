package main

func main() {

}

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		res[i] = make([]int, n-2)
	}
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			maxValue := grid[i][j]
			for a := i - 1; a <= i+1; a++ {
				for b := j - 1; b <= j+1; b++ {
					maxValue = max(maxValue, grid[a][b])
				}
			}
			res[i-1][j-1] = maxValue
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
