package main

import "fmt"

func main() {
	fmt.Println(minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
}

var m []bool

func minimumJumps(forbidden []int, a int, b int, x int) int {
	m = make([]bool, 6001)
	for i := 0; i < len(forbidden); i++ {
		m[forbidden[i]] = true
	}
	return dfs(0, 0, a, b, x)
}

func dfs(index int, dir int, a int, b int, x int) int {
	if index == x {
		return 0
	}
	res := -1
	if index+a < len(m) && m[index+a] == false { // 向前跳+a
		m[index+a] = true
		res = dfs(index+a, 0, a, b, x)
		if res != -1 {
			return res + 1
		}
	}
	if dir == 0 && index-b > 0 && m[index-b] == false { // 向后跳-b
		res = dfs(index-b, 1, a, b, x)
		if res != -1 {
			return res + 1
		}
	}
	return res
}
