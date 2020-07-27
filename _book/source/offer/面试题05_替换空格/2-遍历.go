package main

import (
	"fmt"
)

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

// 剑指offer_面试题05_替换空格
func replaceSpace(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = res + "%20"
		} else {
			res = res + string(s[i])
		}
	}
	return res
}
