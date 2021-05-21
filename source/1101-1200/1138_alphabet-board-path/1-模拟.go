package main

import "strings"

func main() {

}

// leetcode1138_字母板上的路径
func alphabetBoardPath(target string) string {
	res := ""
	x, y := 0, 0
	for i := 0; i < len(target); i++ {
		newX := int(target[i]-'a') / 5
		newY := int(target[i]-'a') % 5
		if x > newX {
			res = res + strings.Repeat("U", x-newX) // 优先向上：从Z往上
		}
		if y > newY {
			res = res + strings.Repeat("L", y-newY) // 优先向左：往Z走
		}
		if y < newY {
			res = res + strings.Repeat("R", newY-y) // 向右
		}
		if x < newX {
			res = res + strings.Repeat("D", newX-x) // 向下
		}
		res = res + "!"
		x, y = newX, newY
	}
	return res
}
