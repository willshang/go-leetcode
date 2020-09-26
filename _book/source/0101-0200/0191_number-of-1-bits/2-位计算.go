package main

import "fmt"

func main() {
	// fmt.Println(hammingWeight(3))
	fmt.Println(hammingWeight(23))
}

// 10111(23) & 10110(22) = 10110(22)
// 10110(22) & 10101(21) = 10100(20)
// 10100(20) & 10011(19) = 10000(16)
// 10000(16) & 01111(15) = 00000(0)
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
