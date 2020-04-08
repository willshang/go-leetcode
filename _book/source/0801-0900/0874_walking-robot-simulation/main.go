package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}

	fmt.Println(robotSim(commands, obstacles))
}

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func robotSim(commands []int, obstacles [][]int) int {
	i := 0
	x := 0
	y := 0
	maxdis := 0
	ob := map[string]bool{}

	for _, v := range obstacles {

		key := ""
		for _, k := range v {
			key = key + strconv.Itoa(k) + ","
		}
		ob[strings.Trim(key, ",")] = true
	}

	for _, v := range commands {
		if v == -2 {
			i = (i + 3) % 4 //左转
		} else if v == -1 {
			i = (i + 1) % 4 //右转
		} else {
			for v > 0 {
				ddx := x + dx[i]
				ddy := y + dy[i]
				tp := strconv.Itoa(ddx) + "," + strconv.Itoa(ddy)
				if _, ok := ob[tp]; ok {
					break
				} else {
					x = ddx
					y = ddy
					maxdis = max(maxdis, x*x+y*y)
				}
				v--
			}
		}
	}
	return maxdis
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*var dxs = []int{0,1,0,-1}
var dys = []int{1,0,-1,0}
//上、右、下、左
func robotSim(commands []int, obstacles [][]int) int {
	isBlocked := make(map[int]bool,10000)
	for _, o := range obstacles{
		i, j := o[0],o[1]
		isBlocked[encode(i,j)] = true
	}

	x,y, res := 0,0,0
	index := 0

	for _, c := range commands{
		switch  {
		case c == -2:
			index--
		case c == -1:
			index++
		default:
			if index < 0{
				index +=  1<<63 - 4
			}

			index = index % 4
			dx, dy := dxs[index], dys[index]
			for c > 0 && !isBlocked[encode(x+dx,y+dy)]{
				c--
				x = x + dx
				y = y + dy
			}
			res = max(res,x*x+y*y)
		}
	}
	return res
}

func encode(x,y int)int  {
	x &= 0xFFFF
	y &= 0xFFFF

	return x << 16 | y
}

func max(a, b int)int  {
	if a > b{
		return a
	}
	return b
}*/
