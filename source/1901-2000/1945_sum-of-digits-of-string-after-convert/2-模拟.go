package main

import "fmt"

func main() {
	fmt.Println(getLucky("iiii", 1))
}

// leetcode1945_字符串转化后的各位数字之和
func getLucky(s string, k int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		value := int(s[i]-'a') + 1
		sum = sum + (value/10 + value%10)
	}
	res := sum
	for i := 1; i <= k-1; i++ {
		sum = res
		res = 0
		for sum > 0 {
			res = res + sum%10
			sum = sum / 10
		}
	}
	return res
}
