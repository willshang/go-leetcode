package main

import "fmt"

func main() {
	//fmt.Println(-1 / -2)
	//fmt.Println(-(-1 >> 1))
	fmt.Println(baseNeg2(3))
}

// leetcode1017_负二进制转换
func baseNeg2(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	for n != 0 {
		if n%2 == 0 { // 偶数
			res = "0" + res
		} else { // 奇数
			res = "1" + res
		}
		// 3 = 111
		// -2做法：-1 / -2 = 0
		// 位做法：-(-1 >> 1) = 1
		n = -(n >> 1) // 除以-2
	}
	return res
}
