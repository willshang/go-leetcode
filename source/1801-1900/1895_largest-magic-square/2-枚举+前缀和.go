package main

func main() {

}

// leetcode1895_最大的幻方
func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	rowArr := make([][]int, n+1)   // 行前缀和
	colArr := make([][]int, n+1)   // 列前缀和
	leftArr := make([][]int, n+1)  // 对角线
	rightArr := make([][]int, n+1) // 对角线
	for i := 0; i <= n; i++ {
		rowArr[i] = make([]int, m+1)
		colArr[i] = make([]int, m+1)
		leftArr[i] = make([]int, m+1)
		rightArr[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rowArr[i][j+1] = rowArr[i][j] + grid[i][j]
			colArr[i+1][j] = colArr[i][j] + grid[i][j]
			leftArr[i+1][j+1] = leftArr[i][j] + grid[i][j]
			rightArr[i+1][j] = rightArr[i][j+1] + grid[i][j]
		}
	}
	for length := min(n, m); length >= 2; length-- { // 枚举边长
		for i := 0; i+length <= n; i++ {
			for j := 0; j+length <= m; j++ {
				target := leftArr[i+length][j+length] - leftArr[i][j]
				if rightArr[i+length][j]-rightArr[i][j+length] != target {
					continue
				}
				flag := true
				for k := i; k < i+length; k++ { // 检查行
					temp := rowArr[k][j+length] - rowArr[k][j]
					if temp != target {
						flag = false
						break
					}
				}
				if flag == false {
					continue
				}
				for k := j; k < j+length; k++ { // 检查列
					temp := colArr[i+length][k] - colArr[i][k]
					if temp != target {
						flag = false
						break
					}
				}
				if flag == true {
					return length
				}
			}
		}
	}
	return 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
