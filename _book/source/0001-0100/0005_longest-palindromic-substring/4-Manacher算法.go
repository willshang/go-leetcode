package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	str := add(s)
	length := len(str)
	max := 1
	begin := 0
	for i := 0; i < length; i++ {
		curLength := search(str, i)
		if curLength > max {
			max = curLength
			begin = (i - max) / 2
		}
	}
	return s[begin : begin+max]
}

func search(s string, center int) int {
	i := center - 1
	j := center + 1
	step := 0
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		step++
	}
	return step
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
