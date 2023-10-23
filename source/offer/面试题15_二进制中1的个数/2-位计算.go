package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(15))
}

// 剑指offer_面试题15_二进制中1的个数
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
