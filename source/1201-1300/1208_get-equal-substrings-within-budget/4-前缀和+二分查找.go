package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
}

func equalSubstring(s string, t string, maxCost int) int {
	res := 0
	arr := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		arr[i+1] = arr[i] + getValue(s[i], t[i])
	}
	for i := 1; i <= len(s); i++ {
		index := sort.SearchInts(arr[:i], arr[i]-maxCost)
		if i-index > res {
			res = i - index
		}
	}
	return res
}

func getValue(a, b byte) int {
	res := int(a) - int(b)
	if res < 0 {
		return -res
	}
	return res
}
