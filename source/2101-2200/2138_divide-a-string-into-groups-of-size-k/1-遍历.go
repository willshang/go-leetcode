package main

import "strings"

func main() {

}

// leetcode2138_将字符串拆分为若干长度为k的组
func divideString(s string, k int, fill byte) []string {
	n := len(s)
	res := make([]string, 0)
	for i := 0; i < n; i = i + k {
		if i+k <= n {
			res = append(res, s[i:i+k])
		} else {
			res = append(res, s[i:]+strings.Repeat(string(fill), k-(n-i)))
		}
	}
	return res
}
