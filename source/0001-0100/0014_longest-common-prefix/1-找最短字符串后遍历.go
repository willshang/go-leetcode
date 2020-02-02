package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	short := strs[0]
	for _, s := range strs {
		if len(short) > len(s) {
			short = s
		}
	}

	for i := range short {
		shortest := short[:i+1]
		for _, str := range strs {
			if strings.Index(str, shortest) != 0 {
				return short[:i]
			}
		}
	}
	return short
}
