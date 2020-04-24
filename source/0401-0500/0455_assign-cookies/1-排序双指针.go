package main

import (
	"fmt"
	"sort"
)

func main() {
	stu := []int{1, 2}
	cookie := []int{1, 2, 3}
	fmt.Println(findContentChildren(stu, cookie))
}

// leetcode455_分发饼干
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
