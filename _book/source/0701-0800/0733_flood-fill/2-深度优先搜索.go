package main

import "fmt"

func main() {
	arr := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(arr, 1, 1, 2))
}

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if sr < 0 || sc < 0 || sr >= len(image) ||
		sc >= len(image[sr]) || image[sr][sc] == newColor {
		return image
	}
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	for i := 0; i < 4; i++ {
		x := sr + dx[i]
		y := sc + dy[i]
		if 0 <= x && x < len(image) && 0 <= y && y < len(image[x]) &&
			image[x][y] == oldColor {
			floodFill(image, x, y, newColor)
		}
	}
	return image
}
