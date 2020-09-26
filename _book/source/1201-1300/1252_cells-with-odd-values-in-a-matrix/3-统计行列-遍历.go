package main

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	//fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
}

// leetcode1252_奇数值单元格的数目
func oddCells(n int, m int, indices [][]int) int {
	rows := make([]int, n)
	cols := make([]int, m)
	for i := 0; i < len(indices); i++ {
		rows[indices[i][0]]++
		cols[indices[i][1]]++
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (rows[i]+cols[j])%2 == 1 {
				res++
			}
		}
	}
	return res
}
