package main

import "math/bits"

func main() {

}

// leetcode1680_连接连续二进制数字
func concatenatedBinary(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		length := bits.Len(uint(i))
		res = (res<<length + i) % 1000000007
	}
	return res
}
