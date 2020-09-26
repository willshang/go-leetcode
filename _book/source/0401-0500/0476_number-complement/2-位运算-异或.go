package main

import "fmt"

func main() {
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(8))
}

// 5二进制为101，则补数等于111^101=010
func findComplement(num int) int {
	temp := num
	res := 0
	for temp > 0 {
		temp = temp >> 1
		res = res << 1
		res++
	}
	return res ^ num
}
