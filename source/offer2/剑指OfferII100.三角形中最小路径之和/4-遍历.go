package main

import "fmt"

func main() {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(arr))
}

// 剑指OfferII100.三角形中最小路径之和
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
