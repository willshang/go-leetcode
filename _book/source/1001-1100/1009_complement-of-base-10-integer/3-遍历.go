package main

import "fmt"

func main() {
	fmt.Println(bitwiseComplement(5))
	fmt.Println(bitwiseComplement(0))
}

func bitwiseComplement(N int) int {
	res := 0
	if N == 0 {
		return 1
	}
	if N == 1 {
		return 0
	}
	exp := 1
	for N > 0 {
		if N%2 == 0 {
			res = res + exp
		}
		exp = exp * 2
		N = N / 2
	}
	return res
}
