package main

import "fmt"

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBits(4294967293))
}

// leetcode 190 颠倒二进制位
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		last := num & 1               // 取最后一位
		result = (result << 1) + last // 前移
		num = num >> 1
	}
	return result
}
