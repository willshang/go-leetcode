package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(checkIfCanBreak("abc", "xya"))
	//fmt.Println(checkIfCanBreak("leetcodee", "interview"))
	fmt.Println(checkIfCanBreak("abe", "acd"))
}

// leetcode1433_检查一个字符串是否可以打破另一个字符串
func checkIfCanBreak(s1 string, s2 string) bool {
	arr1 := []byte(s1)
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	arr2 := []byte(s2)
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	s1 = string(arr1)
	s2 = string(arr2)
	return compare(s1, s2) || compare(s2, s1)
}

func compare(s1 string, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] < s2[i] {
			return false
		}
	}
	return true
}
