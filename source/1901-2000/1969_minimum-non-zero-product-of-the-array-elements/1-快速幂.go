package main

import "fmt"

func main() {
	fmt.Println(minNonZeroProduct(31))
}

// leetcode1969_数组元素的最小非零乘积
var mod = 1000000007

func minNonZeroProduct(p int) int {
	// 2个数x,y交换后，x+y的和是不变的
	// 和不变，要求非0乘积xy最小 => x=1,y=2^p-1
	// 最后：1个2^p-1，2^(p-1)-1个1和2^p-2
	a := (1<<p - 1) % mod
	b := (1<<p - 2) % mod
	c := 1<<(p-1) - 1 // 指数不mod
	return a * mypow(b, c) % mod
}

func mypow(a int, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n = n / 2
	}
	return res
}
