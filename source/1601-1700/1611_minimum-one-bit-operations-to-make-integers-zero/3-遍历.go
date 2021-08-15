package main

import "math/bits"

func main() {

}

// leetcode1611_使整数变为0的最少操作次数
func minimumOneBitOperations(n int) int {
	// 依次将高位的1翻转为0
	// 操作2：00..110000..00 => 00..010000..00
	// 操作2+操作1：00..010000..00 => 00..011000..00
	// 操作2：00..011000..00 => 00..001000..00
	res := 0
	if n == 0 {
		return 0
	}
	length := bits.Len(uint(n))
	flag := 1
	for i := 0; (1 << i) <= n; i++ {
		if (1<<(length-1-i))&n > 0 { // 第length-1-i位>0{
			total := 1<<(length-i) - 1
			res = res + total*flag
			flag = -1 * flag
		}
	}
	return res
}
