package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
}

func removeDuplicates(s string, k int) string {
	if len(s) < k {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		res = res + string(s[i])
		if len(res) >= k && res[len(res)-k:] == strings.Repeat(string(s[i]), k) {
			res = res[:len(res)-k]
		}
	}
	return res
}
