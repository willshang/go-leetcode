package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}

// 顺时针：上右下左
// 本题：右、下、左、上
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	res := make([][]int, 0)
	total := rows * cols
	x, y := rStart, cStart
	res = append(res, []int{x, y})
	index := 1
	if total == 1 {
		return res
	}
	for k := 1; k < 2*(rows+cols); k = k + 2 {
		for i := 0; i < 4; i++ {
			step := k + i/2 // 本次移动的步数，分别是2次的倍数
			for j := 0; j < step; j++ {
				x = x + dx[i]
				y = y + dy[i]
				if 0 <= x && x < rows && 0 <= y && y < cols {
					res = append(res, []int{x, y})
					index++
					if index == total {
						return res
					}
				}
			}

		}
	}
	return res
}
