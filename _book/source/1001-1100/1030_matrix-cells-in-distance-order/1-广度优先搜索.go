package main

import "fmt"

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0)
	visited := make([][]bool, R)
	for i := 0; i < R; i++ {
		visited[i] = make([]bool, C)
	}
	list := make([][]int, 0)
	list = append(list, []int{r0, c0})
	visited[r0][c0] = true
	for len(list) > 0 {
		x1, y1 := list[0][0], list[0][1]
		res = append(res, []int{x1, y1})
		list = list[1:]
		for i := 0; i < 4; i++ {
			x := x1 + dx[i]
			y := y1 + dy[i]
			if (0 <= x && x < R && 0 <= y && y < C) && visited[x][y] == false {
				visited[x][y] = true
				list = append(list, []int{x, y})
			}
		}
	}
	return res
}
