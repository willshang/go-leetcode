package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(arr))
}

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			if grid[i+1][j+1] != 5 {
				continue
			}
			if available(i, j, grid) {
				res++
			}
		}
	}
	return res
}

var m = map[string]bool{
	"816357492": true,
	"834159672": true,
	"618753294": true,
	"672159834": true,
	"492357816": true,
	"438951276": true,
	"294753618": true,
	"276951438": true,
}

func available(x, y int, g [][]int) bool {
	str := ""
	for i := x; i <= x+2; i++ {
		for j := y; j <= y+2; j++ {
			str = str + strconv.Itoa(g[i][j])
		}
	}
	if m[str] {
		return true
	}
	return false
}
