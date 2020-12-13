package main

func main() {

}

func findSquare(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	n, m := len(matrix), len(matrix[0])
	res := 0
	x := 0
	y := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				if 1 > res {
					res = 1
					x, y = i, j
				}
				minLength := min(n-i, m-j)
				for k := 1; k < minLength; k++ {
					flag := true
					if matrix[i+k][j+k] == '1' {
						break
					}
					for l := 0; l < k; l++ {
						if matrix[i+k][j+l] == '1' || matrix[i+l][j+k] == '1' {
							flag = false
							break
						}
					}
					if flag == true {
						if k+1 > res {
							res = k + 1
							x, y = i, j
						}
					} else {
						break
					}
				}
			}
		}
	}
	return []int{res, x, y}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
