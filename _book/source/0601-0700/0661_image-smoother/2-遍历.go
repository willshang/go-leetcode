package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(imageSmoother(arr))
}

// leetcode661_图片平滑器
func imageSmoother(M [][]int) [][]int {
	res := make([][]int, len(M))
	for i := range res {
		res[i] = make([]int, len(M[0]))
		for j := range res[i] {
			value, count := 0, 0
			for r := i - 1; r <= i+1; r++ {
				for c := j - 1; c <= j+1; c++ {
					if 0 <= r && r < len(M) && 0 <= c && c < len(M[0]) {
						value = value + M[r][c]
						count++
					}
				}
			}
			res[i][j] = value / count
		}
	}
	return res
}
