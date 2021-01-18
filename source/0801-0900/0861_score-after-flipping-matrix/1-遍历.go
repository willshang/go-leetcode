package main

func main() {

}

// leetcode861_翻转矩阵后的得分
func matrixScore(A [][]int) int {
	var res int
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}
	// 翻转行，每行第一个为0则翻转
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			for j := 0; j < len(A[i]); j++ {
				A[i][j] = 1 - A[i][j]
			}
		}
	}
	// 翻转列,每列1的数量大于0则翻转
	for j := 0; j < len(A[0]); j++ {
		a, b := 0, 0
		for i := 0; i < len(A); i++ {
			if A[i][j] == 0 {
				a++
			} else {
				b++
			}
		}
		if a <= b {
			continue
		}
		for i := 0; i < len(A); i++ {
			A[i][j] = 1 - A[i][j]
		}
	}
	for i := 0; i < len(A); i++ {
		sum := 0
		for j := 0; j < len(A[i]); j++ {
			sum = sum*2 + A[i][j]
		}
		res = res + sum
	}
	return res
}
