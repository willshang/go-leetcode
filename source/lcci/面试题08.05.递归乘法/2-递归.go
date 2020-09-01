package main

import "fmt"

func main() {
	fmt.Println(multiply(3, 4))
}

// 程序员面试金典08.05_递归乘法
func multiply(A int, B int) int {
	if B == 0 {
		return 0
	}
	if B == 1 {
		return A
	}
	if B%2 == 1 {
		return multiply(A<<1, B>>1) + A
	}
	return multiply(A<<1, B>>1)
}
