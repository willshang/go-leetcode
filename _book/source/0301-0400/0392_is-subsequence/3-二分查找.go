package main

import (
	"fmt"
)

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdcc"))
}

func isSubsequence(s string, t string) bool {
	m := make(map[uint8][]int)
	for i := 0; i < len(t); i++ {
		value := t[i] - 'a'
		if m[value] == nil {
			m[value] = make([]int, 0)
		}
		m[value] = append(m[value], i)
	}
	prev := -1
	for i := 0; i < len(s); i++ {
		value := s[i] - 'a'
		left := 0
		right := len(m[value]) - 1
		if len(m[value]) == 0 {
			return false
		}
		for left < right {
			mid := left + (right-left)/2
			if m[value][mid] > prev {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left > right || m[value][left] <= prev {
			return false
		}
		prev = m[value][left]
	}
	return true
}
