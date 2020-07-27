package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isIsomorphic("foo", "baa")) // True
}

//  strings.IndexByte
// 字符c在s中第一次出现的位置，不存在则返回-1
func isIsomorphic(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if strings.IndexByte(s[i+1:], s[i]) != strings.IndexByte(t[i+1:], t[i]) {
			return false
		}
	}
	return true
}
