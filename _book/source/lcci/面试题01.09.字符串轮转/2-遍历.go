package main

import (
	"fmt"
)

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
}

// 程序员面试金典01.09_字符串轮转
func isFlipedString(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		s1 = s1[1:] + string(s1[0])
		if s1 == s2 {
			return true
		}
	}
	return false
}
