package main

import "fmt"

func main() {
	fmt.Println(reconstructMatrix(2, 1, []int{1, 1, 1}))
}

// leetcode1253_重构2行二进制矩阵
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 0)
	up := make([]int, len(colsum))
	down := make([]int, len(colsum))
	upSum := 0
	lowSum := 0
	total := 0
	for i := 0; i < len(colsum); i++ {
		total = total + colsum[i]
		if colsum[i] == 2 { // 2=>各填充1
			up[i] = 1
			down[i] = 1
			upSum++
			lowSum++
		}
	}
	if upSum > upper || lowSum > lower || total != upper+lower {
		return nil
	}
	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 1 {
			if upSum < upper {
				up[i] = 1
				upSum++
			} else {
				down[i] = 1
			}
		}
	}
	res = append(res, up, down)
	return res
}
