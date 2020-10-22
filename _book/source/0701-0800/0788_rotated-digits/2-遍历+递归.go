package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotatedDigits(10))
}

// 好数的标准就是包含2，5，6，9中的至少一个（保证翻转后数值不同）
// 并且不能包含3,4,7
// 0, 1, 8可以有
func rotatedDigits(N int) int {
	count := 0
	for i := 2; i <= N; i++ {
		if isValid(i, false) {
			count++
		}
	}
	return count
}

func isValid(n int, flag bool) bool {
	if n == 0 {
		return flag
	}
	switch n % 10 {
	case 3, 4, 7:
		return false
	case 0, 1, 8:
		return isValid(n/10, flag)
	case 2, 5, 6, 9:
		return isValid(n/10, true)
	}
	return false
}
