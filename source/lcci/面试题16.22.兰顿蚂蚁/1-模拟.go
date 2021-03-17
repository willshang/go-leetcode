package main

import "fmt"

func main() {
	fmt.Println(printKMoves(1))
}

// 程序员面试金典16.22_兰顿蚂蚁
func printKMoves(K int) []string {
	var dirArr = []byte{'R', 'D', 'L', 'U'}
	var dx = []int{1, 0, -1, 0}
	var dy = []int{0, -1, 0, 1}
	dir := 0 // 向右
	x, y := 0, 0
	left, right := 0, 0
	up, down := 0, 0
	m := make(map[[2]int]int) // 1黑色，0白色
	for i := 0; i < K; i++ {
		if m[[2]int{x, y}] == 1 { // 变方向
			dir = (dir + 3) % 4 // 逆时针
		} else {
			dir = (dir + 1) % 4 // 顺时针
		}
		m[[2]int{x, y}] = 1 - m[[2]int{x, y}]
		x = x + dx[dir]
		y = y + dy[dir]
		left = min(left, x)
		right = max(right, x)
		down = min(down, y)
		up = max(up, y)
	}
	w := right - left + 1
	h := up - down + 1
	res := make([]string, 0)
	for i := 0; i < h; i++ {
		arr := make([]byte, w)
		for j := 0; j < w; j++ {
			newX := j + left
			newY := up - i
			arr[j] = '_'
			if v, ok := m[[2]int{newX, newY}]; ok && v == 1 {
				arr[j] = 'X'
			}
			if newX == x && newY == y {
				arr[j] = dirArr[dir]
			}
		}
		res = append(res, string(arr))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
