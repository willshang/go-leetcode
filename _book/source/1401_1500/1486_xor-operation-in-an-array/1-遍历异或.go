package main

import "fmt"

func main() {
	fmt.Println(xorOperation(5, 0))
}

// leetcode1486_数组异或操作
func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res
}
