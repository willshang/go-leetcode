package main

import "fmt"

func main() {
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(8))
}

// leetcode476_数字的补数
func findComplement(num int) int {
	// 5二进制为101，则补数等于111-101=010=1000-101-1
	temp := 1
	for num >= temp {
		temp = temp << 1
	}
	return temp - 1 - num
}
