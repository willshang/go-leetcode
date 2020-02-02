package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(3))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	strs := countAndSay(n - 1)

	result := make([]byte, 0, len(strs)*2)
	i, j := 0, 1
	for i < len(strs) {
		for j < len(strs) && strs[i] == strs[j] {
			j++
		}
		// 几个几
		result = append(result, byte(j-i+'0'))
		result = append(result, strs[i])
		i = j
	}
	return string(result)
}
