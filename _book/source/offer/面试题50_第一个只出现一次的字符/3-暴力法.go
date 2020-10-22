package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
	fmt.Println(firstUniqChar(""))
}

func firstUniqChar(s string) byte {
	res := byte(' ')
	for i := 0; i < len(s); i++ {
		flag := true
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j {
				flag = false
				break
			}
		}
		if flag {
			return s[i]
		}
	}
	return res
}
