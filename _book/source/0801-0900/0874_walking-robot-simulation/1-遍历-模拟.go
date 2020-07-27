package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}

// leetcode874_模拟行走机器人
// 上、右、下、左
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func robotSim(commands []int, obstacles [][]int) int {
	i := 0 // 方向， 0上， 1右， 2下， 3左
	x := 0
	y := 0
	res := 0
	m := map[string]bool{}
	for _, v := range obstacles {
		str := strconv.Itoa(v[0]) + "," + strconv.Itoa(v[1])
		m[str] = true
	}
	for _, v := range commands {
		if v == -2 {
			i = (i + 3) % 4 // 左转
		} else if v == -1 {
			i = (i + 1) % 4 // 右转
		} else {
			for v > 0 {
				ddx := x + dx[i]
				ddy := y + dy[i]
				tp := strconv.Itoa(ddx) + "," + strconv.Itoa(ddy)
				if _, ok := m[tp]; ok {
					// 有障碍物，停止
					break
				} else {
					x = ddx
					y = ddy
					if x*x+y*y > res {
						res = x*x + y*y
					}
				}
				v--
			}
		}
	}
	return res
}
