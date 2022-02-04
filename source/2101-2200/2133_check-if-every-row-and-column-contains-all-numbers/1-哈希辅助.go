package main

func main() {

}

// leetcode2133_检查是否每一行每一列都包含全部整数
func checkValid(matrix [][]int) bool {
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		visited := make(map[int]bool)
		for j := 0; j < m; j++ {
			if visited[matrix[i][j]] == true {
				return false
			}
			visited[matrix[i][j]] = true
		}
	}
	for j := 0; j < m; j++ {
		visited := make(map[int]bool)
		for i := 0; i < n; i++ {
			if visited[matrix[i][j]] == true {
				return false
			}
			visited[matrix[i][j]] = true
		}
	}
	return true
}
