package main

import "fmt"

func main() {
	a := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	}
	b := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 1},
	}
	fmt.Println(largestOverlap(a, b))
}

func largestOverlap(img1 [][]int, img2 [][]int) int {
	res := 0
	n := len(img1)
	A := make([][2]int, 0)
	B := make([][2]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img1[i][j] == 1 {
				A = append(A, [2]int{i, j})
			}
			if img2[i][j] == 1 {
				B = append(B, [2]int{i, j})
			}
		}
	}
	m := make(map[[2]int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			a := B[j][0] - A[i][0]
			b := B[j][1] - A[i][1]
			m[[2]int{a, b}]++
			res = max(res, m[[2]int{a, b}])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
