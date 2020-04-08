package main

import "fmt"

func main() {
	res := generate(5)
	for _, v := range res {
		fmt.Println(v)
	}
}

// leetcode 118 杨辉三角
func generate(numRows int) [][]int {
	var result [][]int
	for i := 0; i < numRows; i++ {
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
	return result
}
