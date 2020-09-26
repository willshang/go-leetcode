package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(robot("URR", [][]int{{2, 2}}, 3, 2))
}

// leetcode_lcp03_机器人大冒险
func robot(command string, obstacles [][]int, x int, y int) bool {
	if judge(command, x, y) == false {
		return false
	}
	for _, node := range obstacles {
		if x >= node[0] && y >= node[1] && judge(command, node[0], node[1]) {
			return false
		}
	}
	return true
}

func judge(command string, x, y int) bool {
	u := strings.Count(command, "U")
	r := strings.Count(command, "R")
	times := (x + y) / len(command)
	last := command[:(x+y)%len(command)]
	uNum := u*times + strings.Count(last, "U")
	rNum := r*times + strings.Count(last, "R")
	if uNum == y && rNum == x {
		return true
	}
	return false
}
