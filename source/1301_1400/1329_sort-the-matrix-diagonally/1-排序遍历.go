package main

import "sort"

func main() {

}

func diagonalSort(mat [][]int) [][]int {
	m := make(map[int][]int)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			m[i-j] = append(m[i-j], mat[i][j])
		}
	}
	for _, v := range m {
		sort.Ints(v)
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = m[i-j][0]
			m[i-j] = m[i-j][1:]
		}
	}
	return mat
}
