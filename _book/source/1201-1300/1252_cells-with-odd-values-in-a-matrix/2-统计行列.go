package main

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	//fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
}

func oddCells(n int, m int, indices [][]int) int {
	rows := make([]int, n)
	cols := make([]int, m)
	for i := 0; i < len(indices); i++ {
		rows[indices[i][0]]++
		cols[indices[i][1]]++
	}
	numRows := 0
	for i := 0; i < n; i++ {
		if rows[i]%2 == 0 {
			numRows++
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		if cols[i]%2 == 0 {
			res = res + n - numRows
		} else {
			res = res + numRows
		}
	}
	return res
}
