package main

import "fmt"

func main() {
	fmt.Println(getSum(101, 100))
}

/*
0 + 0 = 0
0 + 1 = 1
1 + 0 = 1
1 + 1 = 0（进位 1）
异或的一个重要特性是无进位加法
(a 和 b 的无进位结果) + (a 和 b 的进位结果)
*/

// leetcode371_两整数之和
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
