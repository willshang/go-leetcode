package main

import "fmt"

func main() {
	fmt.Println(checkPalindromeFormation("ulacfd", "jizalu"))
}

// leetcode1616_分割两个字符串得到回文串
func checkPalindromeFormation(a string, b string) bool {
	start := len(a)/2 - 1 // 中间靠左坐标
	c := check(a, a, start)
	if check(a, b, c) == -1 || check(b, a, c) == -1 {
		return true
	}
	c = check(b, b, start)
	if check(a, b, c) == -1 || check(b, a, c) == -1 {
		return true
	}
	return false
}

// 从start往左边走，从n-1-start往右边走
func check(a, b string, start int) int {
	left := start
	right := len(a) - 1 - start
	for left >= 0 {
		if a[left] != b[right] {
			break
		}
		left--
		right++
	}
	return left
}
