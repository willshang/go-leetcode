package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(10))
	fmt.Println(countNumbersWithUniqueDigits(2))
}

// leetcode357_计算各个位数不同的数字个数
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	prev := 1
	for i := 1; i <= 10 && i <= n; i++ {
		res = res + 9*prev
		prev = prev * (10 - i)
	}
	return res
}
