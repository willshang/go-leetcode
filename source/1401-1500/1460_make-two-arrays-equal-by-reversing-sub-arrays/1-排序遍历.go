package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}

// leetcode1460_通过翻转子数组使两个数组相等
func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
