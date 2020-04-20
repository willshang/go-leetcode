package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, str))
}

// leetcode290_单词规律
func wordPattern(pattern string, str string) bool {
	pa := strings.Split(pattern, "")
	sa := strings.Split(str, " ")
	if len(pa) != len(sa) {
		return false
	}
	length := len(pa)
	pMap := make(map[string]string, length)
	sMap := make(map[string]string, length)

	for i := 0; i < length; i++ {
		pStr, ok := pMap[pa[i]]
		sStr, ok1 := sMap[sa[i]]

		if (ok && pStr != sa[i]) || (ok1 && sStr != pa[i]) {
			return false
		} else {
			pMap[pa[i]] = sa[i]
			sMap[sa[i]] = pa[i]
		}
	}
	return true
}
