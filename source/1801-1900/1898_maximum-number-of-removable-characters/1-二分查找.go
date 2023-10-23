package main

import "fmt"

func main() {
	fmt.Println(maximumRemovals("abcacb", "ab", []int{3, 1, 0}))
}

// leetcode1898_可移除字符的最大数目
func maximumRemovals(s string, p string, removable []int) int {
	n := len(removable)
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if judge(s, p, removable, mid) == true {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// leetcode 392.判断子序列
func judge(s string, p string, removable []int, index int) bool {
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
	return i == len(p)
}
