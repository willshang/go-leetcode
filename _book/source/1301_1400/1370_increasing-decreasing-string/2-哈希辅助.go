package main

import "fmt"

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
}

func sortString(s string) string {
	m := make(map[int]int, 26)
	for i := 0; i < len(s); i++ {
		m[int(s[i]-'a')]++
	}
	res := ""
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if m[i] > 0 {
				res = res + string(i+'a')
				m[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if m[i] > 0 {
				res = res + string(i+'a')
				m[i]--
			}
		}
	}
	return res
}
