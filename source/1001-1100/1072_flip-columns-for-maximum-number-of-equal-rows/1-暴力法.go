package main

func main() {

}

// leetcode1072_按列翻转得到最大值等行数
func maxEqualRowsAfterFlips(matrix [][]int) int {
	res := 0
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		count := 0            // 统计与当前行完全一样的行或者完全不同的行的个数，划分为同一组
		arr := make([]int, m) // 翻转当前行的结果
		for j := 0; j < m; j++ {
			arr[j] = 1 - matrix[i][j]
		}
		for k := 0; k < n; k++ {
			if compare(matrix[k], matrix[i]) || compare(matrix[k], arr) {
				count++
			}
		}
		res = max(res, count) // 翻转最大一组的任意一行中的所有0或者1所在列即可
	}
	return res
}

func compare(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
