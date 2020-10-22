package main

import "fmt"

func main() {
	fmt.Println(shortestPalindrome("aabcd"))
}

func shortestPalindrome(s string) string {
	str := add(s)
	index := 0
	for i := len(str) / 2; i <= len(str)/2; i-- {
		j := i
		for ; j > 0; j-- {
			if str[i-j] != str[i+j] {
				break
			}
		}
		if j == 0 {
			index = i
			break
		}
	}
	res := s
	for j := index; j < len(s); j++ {
		res = string(s[j]) + res
	}
	return res
}

func add(s string) string {
	var res []rune
	for _, v := range s {
		res = append(res, '#')
		res = append(res, v)
	}
	res = append(res, '#')
	return string(res)
}
