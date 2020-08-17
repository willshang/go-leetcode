package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isUnique("leetcode"))
}

func isUnique(astr string) bool {
	arr := []byte(astr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
	}
	return true
}
