package main

func main() {

}

func matrixScore(A [][]int) int {
	var res int
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}
	n := len(A)
	m := len(A[0])
	// 首先每行第一个都为1求和，n个长度为m的1x...x
	// 这样保证第一列全为1
	res = res + n*(1<<(m-1))
	for j := 1; j < m; j++ {
		a, b := 0, 0
		for i := 0; i < n; i++ {
			if A[i][0] == 0 && A[i][j] == 0 { // 需要翻转
				b++
			} else if A[i][0] == 1 && A[i][j] == 1 { // 不需要翻转
				b++
			} else {
				a++
			}
		}
		// 1比0多，不需要翻转
		if a <= b {
			res = res + b*(1<<(m-1-j))
		} else {
			res = res + a*(1<<(m-1-j))
		}
	}
	return res
}
