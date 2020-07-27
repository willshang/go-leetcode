package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("LIII"))
}

// leetcode 13 罗马数字转整数
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	last := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := m[s[i]]
		if current < last {
			result = result - current
		} else {
			result = result + current
		}
		last = current
	}
	return result
}
