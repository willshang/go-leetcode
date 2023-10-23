package main

import "fmt"

func main() {
	fmt.Println(maximumRemovals("abcacb", "ab", []int{3, 1, 0}))
}

func maximumRemovals(s string, p string, removable []int) int {
	n := len(removable)
	left, right := 0, n+1
	for left < right {
		mid := left + (right-left)/2
		if judge(s, p, removable, mid) == true {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

// leetcode 392.判断子序列
func judge(s string, p string, removable []int, index int) bool {
	m := make(map[int]bool)
	for i := 0; i < len(removable[:index]); i++ {
		m[removable[i]] = true
	}
	i, j := 0, 0
	for i < len(p) && j < len(s) {
		if p[i] == s[j] && m[j] == false {
			i++
		}
		j++
	}
	return i == len(p)
}
