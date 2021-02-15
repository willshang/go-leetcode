package main

import "fmt"

func main() {
	fmt.Println(maximumGain("abab", 4, 5))
}

// leetcode1717_删除子字符串的最大得分
func maximumGain(s string, x int, y int) int {
	if x > y {
		x, y = y, x
		s = reverse(s)
	}
	res := 0
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'b' {
			if s[i] == 'a' {
				a++
			} else {
				b++
			}
			// 处理ba
			if s[i] == 'a' && b > 0 {
				a--
				b--
				res = res + y
			}
		} else {
			res = res + x*min(a, b) // 处理ab
			a, b = 0, 0
		}
	}
	res = res + x*min(a, b) // 处理ab
	return res
}

func reverse(s string) string {
	arr := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		arr[i], arr[len(s)-1-i] = arr[len(s)-1-i], arr[i]
	}
	return string(arr)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
