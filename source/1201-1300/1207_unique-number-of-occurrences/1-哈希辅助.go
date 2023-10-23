package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}

// leetcode1207_独一无二的出现次数
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	temp := make(map[int]bool)
	for _, v := range m {
		if temp[v] == true {
			return false
		}
		temp[v] = true
	}
	return true
}
