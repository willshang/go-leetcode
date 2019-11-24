package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", ""))
}

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		//haystack[0:0] = ""
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}
