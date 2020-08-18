package main

import "fmt"

func main() {
	fmt.Println(insertBits(0, 31, 0, 4))
	//fmt.Println(insertBits(1024, 19, 2, 6))
}

func insertBits(N int, M int, i int, j int) int {
	res := N
	setZero := 0
	for k := i; k <= j; k++ {
		setZero = setZero | (1 << k)
	}
	res = res&setZero ^ N
	res = res | (M << i)
	return res
}
