package main

func main() {

}

// leetcode1886_判断矩阵经轮转后是否一致
func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		rotate(mat)
		if compare(mat, target) == true {
			return true
		}
	}
	return false
}

// leetcode 48.旋转图像
func rotate(matrix [][]int) {
	n := len(matrix)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[j][n-1-i] = matrix[i][j]
		}
	}
	copy(matrix, arr)
}

func compare(a, b [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
