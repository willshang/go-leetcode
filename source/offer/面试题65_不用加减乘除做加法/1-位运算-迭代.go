package main

import "fmt"

func main() {
	fmt.Println(add(3, 5))
}

// 剑指offer_面试题65_不用加减乘除做加法
// 非进位和：异或运算
// 进位：与运算+左移一位
func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
