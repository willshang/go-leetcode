package main

import "fmt"

func main() {
	fmt.Println(isPathCrossing("NESWW"))
}

// leetcode1496_判断路径是否相交
func isPathCrossing(path string) bool {
	m := make(map[string]bool)
	m["0,0"] = true
	x := 0
	y := 0
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			y = y + 1
		case 'S':
			y = y - 1
		case 'E':
			x = x + 1
		case 'W':
			x = x - 1
		}
		if m[fmt.Sprintf("%d,%d", x, y)] {
			return true
		}
		m[fmt.Sprintf("%d,%d", x, y)] = true
	}
	return false
}
