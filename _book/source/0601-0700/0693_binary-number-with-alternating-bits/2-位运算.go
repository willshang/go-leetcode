package main

import (
	"fmt"
)

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(4))
}

/*
示例1：
1. n=1010
2. n>>1=101
3. n=n^(n>>1)=1010^101=1111
4. n&(n+1)=1111&(10000)=0

示例2:
1. n=101
2. n>>1=10
3. n=n^(n>>1)=101^10=111
4. n&(n+1)=111&(1000)=0
*/
// leetcode693_交替位二进制数
func hasAlternatingBits(n int) bool {
	n = n ^ (n >> 1)
	return n&(n+1) == 0
}
