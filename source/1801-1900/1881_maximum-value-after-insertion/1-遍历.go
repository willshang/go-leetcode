package main

import "strings"

func main() {

}

// leetcode1881_插入后的最大值
func maxValue(n string, x int) string {
	if strings.Contains(n, "-") {
		for i := 1; i < len(n); i++ {
			if x < int(n[i]-'0') {
				return n[:i] + string(x+'0') + n[i:]
			}
		}
	} else {
		for i := 0; i < len(n); i++ {
			if x > int(n[i]-'0') {
				return n[:i] + string(x+'0') + n[i:]
			}
		}
	}
	return n + string(x+'0')
}
