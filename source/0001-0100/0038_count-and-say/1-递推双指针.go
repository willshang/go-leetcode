package main

import "fmt"

func main() {
	fmt.Println(countAndSay(8))
}

// leetcode 38 计数
func countAndSay(n int) string {
	strs := []byte{'1'}
	for i := 1; i < n; i++ {
		strs = say(strs)
	}
	return string(strs)
}

func say(strs []byte) []byte {
	result := make([]byte, 0, len(strs)*2)

	i, j := 0, 1
	for i < len(strs) {
		for j < len(strs) && strs[i] == strs[j] {
			j++
		}
		result = append(result, byte(j-i+'0'))
		result = append(result, strs[i])
		i = j
	}
	return result
}
