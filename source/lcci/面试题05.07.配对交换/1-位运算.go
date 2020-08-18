package main

import "fmt"

func main() {
	fmt.Println(exchangeBits(158))
}

// 程序员面试金典05.07_配对交换
func exchangeBits(num int) int {
	// 0x55555555 = 01010101010101010101010101010101 提取偶数位=>左移
	// 0xaaaaaaaa = 10101010101010101010101010101010 提取奇数位=>右移
	a := (num & 0x55555555) << 1
	b := (num & 0xaaaaaaaa) >> 1
	return a | b
}
