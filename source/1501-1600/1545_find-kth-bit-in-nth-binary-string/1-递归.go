package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(findKthBit(4, 15)))
}

// leetcode1545_找出第N个二进制字符串中的第K位
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	mid := 1 << (n - 1)
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	}
	// 下标从1开始，如n=4,k=15=>3,1(2^4-15=16-15=1)
	return change(findKthBit(n-1, (1<<n)-k))
}

func change(char byte) byte {
	if char == '0' {
		return '1'
	}
	return '0'
}
