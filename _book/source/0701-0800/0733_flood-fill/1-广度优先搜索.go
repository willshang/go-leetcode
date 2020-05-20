package main

import "fmt"

func main() {
	arr := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(arr, 1, 1, 2))
}

// leetcode733_图像渲染
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m, n := len(image), len(image[0])
	list := make([][]int, 1)
	list[0] = []int{sr, sc}

	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		image[node[0]][node[1]] = newColor
		for i := 0; i < 4; i++ {
			x := node[0] + dx[i]
			y := node[1] + dy[i]
			if 0 <= x && x < m && 0 <= y && y < n &&
				image[x][y] == oldColor {
				list = append(list, []int{x, y})
			}
		}
	}
	return image
}
