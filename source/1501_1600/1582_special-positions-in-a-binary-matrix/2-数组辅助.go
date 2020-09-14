package main

func main() {

}

// leetcode1582_二进制矩阵中的特殊位置
func numSpecial(mat [][]int) int {
	res := 0
	rows, cols := make([]int, len(mat)), make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			rows[i] = rows[i] + mat[i][j]
			cols[j] = cols[j] + mat[i][j]
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				res++
			}
		}
	}
	return res
}
