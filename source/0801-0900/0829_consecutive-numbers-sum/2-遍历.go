package main

import "fmt"

func main() {
	fmt.Println(consecutiveNumbersSum(5))
}

func consecutiveNumbersSum(N int) int {
	res := 1
	for i := 1; ; i++ {
		N = N - i
		if N > 0 {
			if N%(i+1) == 0 { // 划分为i+1个数
				res++
			}
		} else {
			break
		}
	}
	return res
}
