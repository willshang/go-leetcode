package main

func main() {

}

// leetcode1992_找到所有的农场组
func findFarmland(land [][]int) [][]int {
	res := make([][]int, 0)
	n, m := len(land), len(land[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 0 {
				continue
			}
			if (0 < i && land[i-1][j] == 1) || // 左边或者上边为1
				(0 < j && land[i][j-1] == 1) {
				continue
			}
			var a, b, c, d int
			a, b = i, j
			for c = i; c+1 < n && land[c+1][j] == 1; c++ { // 往下遍历
			}
			for d = j; d+1 < m && land[i][d+1] == 1; d++ { // 往右遍历
			}
			res = append(res, []int{a, b, c, d})
		}
	}
	return res
}
