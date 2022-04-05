package main

import "fmt"

func main() {
	fmt.Println(countCollisions("LL"))
}

// leetcode2211_统计道路上的碰撞次数
func countCollisions(directions string) int {
	n := len(directions)
	i, j := 0, n-1
	for i < n && directions[i] == 'L' {
		i++
	}
	for j >= 0 && directions[j] == 'R' {
		j--
	}
	count := 0
	for k := i; k <= j; k++ {
		if directions[k] == 'S' {
			count++
		}
	}
	return j - i + 1 - count
}
