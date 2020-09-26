package main

import (
	"fmt"
)

func main() {
	str := "LL"
	fmt.Println(judgeCircle(str))
}

// leetcode657_机器人能否返回原点
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := range moves {
		switch moves[i] {
		case 'U':
			y = y + 1
		case 'D':
			y = y - 1
		case 'L':
			x = x - 1
		case 'R':
			x = x + 1
		}
	}
	return x == 0 && y == 0
}
