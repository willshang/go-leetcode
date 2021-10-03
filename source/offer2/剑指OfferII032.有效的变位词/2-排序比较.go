package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	sArr := make([]int, len(s))
	tArr := make([]int, len(t))
	for i := 0; i < len(s); i++ {
		sArr[i] = int(s[i] - 'a')
		tArr[i] = int(t[i] - 'a')
	}
	sort.Ints(sArr)
	sort.Ints(tArr)
	for i := 0; i < len(s); i++ {
		if sArr[i] != tArr[i] {
			return false
		}
	}
	return true
}
