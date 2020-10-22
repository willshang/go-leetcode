package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}

// leetcode201_数字范围按位与
func rangeBitwiseAnd(m int, n int) int {
	count := 0
	// 找m,n的32位二进制，前面相同的位数，然后后面添0
	for m != n {
		count++
		// 同时右移去除末尾1位
		m = m >> 1
		n = n >> 1
	}
	return m << count
}
