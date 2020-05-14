package main

import (
	"fmt"
	"strings"
)

func main() {
	A := "abcd"
	B := "cdabcdab"
	fmt.Println(repeatedStringMatch(A, B))
	fmt.Println(repeatedStringMatch("abc", "cabcabca"))
}

// leetcode686_重复叠加字符串匹配
func repeatedStringMatch(A string, B string) int {
	times := len(B) / len(A)
	// 要确保B是A的子串，就要最少重复len(B)/len(A)次A次，最多len(B)/len(A)+2次
	// 如长度为 len(B) = 6, len(A) = 3,至少重复2次
	// 长度为len(B) = 7, len(A) = 3, 至少重复3次
	// 另外如B="cabcabca", A="abc",需要重复4次
	for i := times; i <= times+2; i++ {
		if strings.Contains(strings.Repeat(A, i), B) {
			return i
		}
	}
	return -1
}
