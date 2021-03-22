package main

import "fmt"

func main() {
	// fmt.Println(totalMoney(4))
	fmt.Println(totalMoney(10))
}

// leetcode1716_计算力扣银行的钱
func totalMoney(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		a, b := i/7, i%7+1
		res = res + a + b
	}
	return res
}
