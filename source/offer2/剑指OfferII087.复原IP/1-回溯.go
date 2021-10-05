package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

// 剑指OfferII087.复原IP
var res []string

func restoreIpAddresses(s string) []string {
	res = make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	dfs(s, make([]string, 0), 0)
	return res
}

func dfs(s string, arr []string, level int) {
	if level == 4 {
		if len(s) == 0 {
			str := strings.Join(arr, ".")
			res = append(res, str)
		}
		return
	}
	for i := 1; i <= 3; i++ {
		if i <= len(s) {
			value, _ := strconv.Atoi(s[:i])
			if value <= 255 {
				str := s[i:]
				dfs(str, append(arr, s[:i]), level+1)
			}
			if value == 0 {
				// 避免出现001,01这种情况
				break
			}
		}
	}
}
