package main

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
}

func oddCells(n int, m int, indices [][]int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	for i := 0; i < len(indices); i++ {
		r := indices[i][0]
		c := indices[i][1]
		for j := 0; j < m; j++ {
			arr[r][j]++
		}
		for j := 0; j < n; j++ {
			arr[j][c]++
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j]%2 == 1 {
				res++
			}
		}
	}
	return res
}
