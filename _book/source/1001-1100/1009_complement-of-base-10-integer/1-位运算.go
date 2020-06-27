package main

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(0))
}

// leetcode1009_十进制整数的反码
/*
101+010=1000=111+1
*/
func bitwiseComplement(N int) int {
	temp := 2
	for N >= temp {
		temp = temp << 1
	}
	return temp - 1 - N
}
