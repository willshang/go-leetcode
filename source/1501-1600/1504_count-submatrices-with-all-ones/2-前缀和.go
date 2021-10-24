package main

func main() {

}

// leetcode1504_统计全1子矩形
func numSubmat(mat [][]int) int {
	res := 0
	n, m := len(mat), len(mat[0])
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				arr[i][j] = mat[i][j]
			} else if j > 0 && mat[i][j] == 1 {
				arr[i][j] = arr[i][j-1] + 1 // 到左边连续1的个数
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			target := arr[i][j]
			for k := i; k >= 0; k-- {
				if target > arr[k][j] {
					target = arr[k][j]
				}
				res = res + target
			}
		}
	}
	return res
}
