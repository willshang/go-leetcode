package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//fmt.Println(string(findTheDifference("abcd", "abcde")))
	fmt.Println(string(findTheDifference("ae", "aae")))
}

func findTheDifference(s string, t string) byte {
	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	sort.Strings(sArr)
	sort.Strings(tArr)
	for i := 0; i < len(sArr); i++ {
		if sArr[i] != tArr[i] {
			return []byte(tArr[i])[0]
		}
	}
	return []byte(tArr[len(tArr)-1])[0]
}
