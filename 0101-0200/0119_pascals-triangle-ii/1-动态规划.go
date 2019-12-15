package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	var result [][]int
	for i := 0; i < rowIndex+1; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			tmp := 1
			if j == 0 || j == i {

			} else {
				tmp = result[i-1][j-1] + result[i-1][j]
			}
			row = append(row, tmp)
		}
		result = append(result, row)
	}
	return result[rowIndex]
}
