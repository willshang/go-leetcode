package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compressString("aabcccccaaa"))
}

// 程序员面试金典01.06_字符串压缩
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	prev := S[0]
	count := 1
	res := ""
	for i := 1; i < len(S); i++ {
		if prev == S[i] {
			count++
		} else {
			res = res + string(prev) + strconv.Itoa(count)
			prev = S[i]
			count = 1
		}
	}
	res = res + string(prev) + strconv.Itoa(count)
	if len(res) >= len(S) {
		return S
	}
	return res
}
