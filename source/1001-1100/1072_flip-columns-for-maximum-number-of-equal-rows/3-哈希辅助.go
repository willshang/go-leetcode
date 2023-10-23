package main

func main() {

}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	res := 0
	n := len(matrix)
	m := len(matrix[0])
	M := make(map[string]int)
	for i := 0; i < n; i++ {
		a := make([]byte, 0)
		if matrix[i][0] == 0 {
			for j := 0; j < m; j++ {
				if matrix[i][j] == 0 {
					a = append(a, '0')
				} else {
					a = append(a, '1')
				}
			}
		} else {
			for j := 0; j < m; j++ {
				if matrix[i][j] == 0 {
					a = append(a, '1')
				} else {
					a = append(a, '0')
				}
			}
		}
		M[string(a)]++
	}
	for _, v := range M {
		res = max(res, v)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
