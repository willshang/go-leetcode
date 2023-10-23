package main

import "fmt"

func main() {
	for i := 0; i <= 22; i++ {
		fmt.Println(i, numberOf2sInRange(i))
	}
}

// 程序员面试金典17.06_2出现的次数
func numberOf2sInRange(n int) int {
	if n <= 0 {
		return 0
	}
	res := 0
	for i := 1; i <= n; i = i * 10 {
		left := n / i
		right := n % i
		res = res + (left+7)/10*i
		if left%10 == 2 {
			res = res + right + 1
		}
	}
	return res
}
