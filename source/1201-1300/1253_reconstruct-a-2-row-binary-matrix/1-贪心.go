package main

import "fmt"

func main() {
	fmt.Println(reconstructMatrix(2, 1, []int{1, 1, 1}))
}

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 0)
	total := 0
	two := 0
	for i := 0; i < len(colsum); i++ {
		total = total + colsum[i]
		if colsum[i] == 2 {
			two++
		}
	}
	if total != upper+lower || two > upper || two > lower {
		return nil
	}
	up := make([]int, len(colsum))
	down := make([]int, len(colsum))
	upper = upper - two // 上面数组单独1的个数
	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 2 { // 2=>各填充1
			up[i] = 1
			down[i] = 1
			lower--
		} else if colsum[i] == 1 {
			if upper > 0 { // 先填充上面数组
				up[i] = 1
				upper--
			} else {
				down[i] = 1
			}
		}
	}
	res = append(res, up, down)
	return res
}
