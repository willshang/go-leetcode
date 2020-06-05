package main

import "fmt"

func main() {
	fmt.Println(lastRemaining(5, 3))
}

// 剑指offer_面试题62_圆圈中最后剩下的数字
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}
