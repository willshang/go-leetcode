package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	str := add(s)
	length := len(str)
	temp := make([]int, length)
	maxRight := 0
	center := 0
	max := 1
	begin := 0
	for i := 0; i < length; i++ {
		if i < maxRight {
			mirror := 2*center - i
			temp[i] = min(maxRight-i, temp[mirror])
		}
		left := i - (1 + temp[i])
		right := i + (1 + temp[i])
		for left >= 0 && right < len(str) && str[left] == str[right] {
			temp[i]++
			left--
			right++
		}
		if i+temp[i] > maxRight {
			maxRight = i + temp[i]
			center = i
		}
		if temp[i] > max {
			max = temp[i]
			begin = (i - max) / 2
		}
	}
	return s[begin : begin+max]
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
