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
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j, res int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}
	return res
}
