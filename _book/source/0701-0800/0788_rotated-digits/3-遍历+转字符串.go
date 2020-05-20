package main

import (
	"fmt"
	"strconv"
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
		if isValid(i) {
			count++
		}
	}
	return count
}

var m = map[byte]byte{
	'0': '0',
	'1': '1',
	'8': '8',
	'2': '5',
	'5': '2',
	'6': '9',
	'9': '6',
}

func isValid(n int) bool {
	str := strconv.Itoa(n)
	newStr := ""
	for i := 0; i < len(str); i++ {
		if value, ok := m[str[i]]; ok {
			newStr = newStr + string(value)
		} else {
			return false
		}
	}
	return newStr != str
}
