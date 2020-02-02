package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(5))
}

// N!有多少个后缀0，即N!有多少个质因数5。
// N!有多少个质因数5，即N可以划分成多少组5个数字一组，
// 加上划分成多少组25个数字一组，加上划分多少组成125个数字一组，等等
// Ans = N/5 + N/(5^2) + N/(5^3) + ...
// leetcode 172 阶乘后的零
func trailingZeroes(n int) int {
	result := 0
	for n >= 5 {
		n = n / 5
		result = result + n
	}
	return result
}
