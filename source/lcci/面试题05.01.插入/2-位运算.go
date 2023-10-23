package main

import "fmt"

func main() {
	fmt.Println(insertBits(0, 31, 0, 4))
	//fmt.Println(insertBits(1024, 19, 2, 6))
}

func insertBits(N int, M int, i int, j int) int {
	for k := i; k <= j; k++ {
		if N&(1<<k) != 0 {
			N = N - 1<<k
		}
	}
	N = N + (M << i)
	return N
}
