package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	for i := 1; i <= 3 && i < len(s)-2; i++ {
		for j := i + 1; j <= i+3 && j < len(s)-1; j++ {
			for k := j + 1; k <= j+3 && k < len(s); k++ {
				if judge(s[:i]) && judge(s[i:j]) &&
					judge(s[j:k]) && judge(s[k:]) {
					res = append(res, s[:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:])
				}
			}
		}
	}
	return res
}

func judge(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	value, _ := strconv.Atoi(s)
	if value > 255 {
		return false
	}
	return true
}
