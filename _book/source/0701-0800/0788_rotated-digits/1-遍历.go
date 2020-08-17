package main

import "fmt"

func main() {
	fmt.Println(rotatedDigits(10))
}

// 好数的标准就是包含2，5，6，9中的至少一个（保证翻转后数值不同）
// 并且不能包含3,4,7
// 0, 1, 8可以有
// leetcode788_旋转数字
func rotatedDigits(N int) int {
	count := 0
	for i := 2; i <= N; i++ {
		if isValid(i) {
			count++
		}
	}
	return count
}

func isValid(n int) bool {
	valid := false
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			valid = true
		case 3, 4, 7:
			return false
		}
		n = n / 10
	}
	return valid
}
