package main

func main() {

}

// leetcode1074_元素和为目标值的子矩阵数量
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res := 0
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[i][j] = matrix[i-1][j-1] + arr[i-1][j] + arr[i][j-1] - arr[i-1][j-1]
		}
	}
	for a := 1; a <= n; a++ { // 上边界
		for b := a; b <= n; b++ { // 下边界
			temp := make(map[int]int)
			temp[0] = 1
			for j := 1; j <= m; j++ {
				value := arr[b][j] - arr[a-1][j]
				res = res + temp[value-target]
				temp[value]++
			}
		}
	}
	return res
}
