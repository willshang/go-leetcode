package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

// leetcode1309_解码字母到整数映射
func freqAlphabets(s string) string {
	res := ""
	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' {
			value, _ := strconv.Atoi(string(s[i : i+2]))
			res = res + string('a'+value-1)
			i = i + 3
		} else {
			value, _ := strconv.Atoi(string(s[i]))
			res = res + string('a'+value-1)
			i = i + 1
		}
	}
	return res
}
