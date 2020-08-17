package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}

// 上、右、下、左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[string]bool, 10000)
	for _, o := range obstacles {
		i, j := o[0], o[1]
		m[encode(i, j)] = true
	}
	x, y, res := 0, 0, 0
	index := 0
	for _, c := range commands {
		index = (index + 4) % 4
		switch {
		case c == -2:
			index--
		case c == -1:
			index++
		default:
			dx1, dy1 := dx[index], dy[index]
			for c > 0 && !m[encode(x+dx1, y+dy1)] {
				c--
				x = x + dx1
				y = y + dy1
			}
			if x*x+y*y > res {
				res = x*x + y*y
			}
		}
	}
	return res
}

func encode(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
