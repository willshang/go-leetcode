package main

func main() {

}

// leetcode1901_找出顶峰元素II
func findPeakGrid(mat [][]int) []int {
	n, m := len(mat), len(mat[0])
	x, y := 0, 0
	for {
		if x < n-1 && mat[x][y] < mat[x+1][y] {
			x++
		} else if 1 <= x && mat[x][y] < mat[x-1][y] {
			x--
		} else if y < m-1 && mat[x][y] < mat[x][y+1] {
			y++
		} else if 1 <= y && mat[x][y] < mat[x][y-1] {
			y--
		} else {
			return []int{x, y}
		}
	}
	return nil
}
