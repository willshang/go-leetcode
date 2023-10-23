package main

import "fmt"

func main() {
	fmt.Println(minSteps(10))
}

// leetcode650_只有两个键的键盘
func minSteps(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res = res + i
			n = n / i
		}
	}
	return res
}
