package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "a##c"
	T := "#a#c"
	fmt.Println(backspaceCompare(S, T))
}

// leetcode844_比较含退格的字符串
func backspaceCompare(S string, T string) bool {
	return check(S) == check(T)
}

func check(str string) string {
	res := make([]string, 0)
	for _, v := range str {
		if string(v) == "#" {
			if len(res) != 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, string(v))
		}
	}
	return strings.Join(res, "")
}
