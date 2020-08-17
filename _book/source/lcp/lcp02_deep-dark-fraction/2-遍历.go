package main

import "fmt"

func main() {
	fmt.Println(fraction([]int{3, 2, 0, 2}))
}

// leetcode_lcp02_分式化简
func fraction(cont []int) []int {
	n, m := 1, cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		n, m = m, cont[i]*m+n
	}
	return []int{m, n}
}
