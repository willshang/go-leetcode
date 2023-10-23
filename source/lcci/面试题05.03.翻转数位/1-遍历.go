package main

import "fmt"

func main() {
	fmt.Println(reverseBits(1775))
}

// 程序员面试金典05.03_翻转数位
func reverseBits(num int) int {
	res := 0
	a, b := 0, 0
	for num != 0 {
		if num%2 == 1 {
			a++
		} else {
			b = a
			a = 0
		}
		res = max(res, a+b)
		num = num / 2
	}
	return res + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
