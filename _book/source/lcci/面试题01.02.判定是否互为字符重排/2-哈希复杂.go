package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}

// 程序员面试金典01.02_判定是否互为字符重排
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		m[s2[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
