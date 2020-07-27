package main

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(0))
}

/*
101^111=010
*/
func bitwiseComplement(N int) int {
	temp := N
	res := 1
	for temp > 1 {
		temp = temp >> 1
		res = res << 1
		res++
	}
	return res ^ N
}
