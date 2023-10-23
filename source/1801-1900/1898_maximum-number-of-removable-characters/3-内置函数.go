package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumRemovals("abcacb", "ab", []int{3, 1, 0}))
}

func maximumRemovals(s string, p string, removable []int) int {
	n := len(removable)
	return sort.Search(n, func(index int) bool {
		m := make(map[int]bool)
		for i := 0; i < len(removable[:index+1]); i++ {
			m[removable[i]] = true
		}
		i, j := 0, 0
		for i < len(p) && j < len(s) {
			if p[i] == s[j] && m[j] == false {
				i++
			}
			j++
		}
		return i != len(p)
	})
}
