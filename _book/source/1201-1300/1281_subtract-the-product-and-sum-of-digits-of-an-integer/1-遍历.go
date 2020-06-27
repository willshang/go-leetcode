package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))
}

// leetcode1281_整数的各位积和之差
func subtractProductAndSum(n int) int {
	sum := 0
	mul := 1
	for n > 0 {
		value := n % 10
		n = n / 10
		sum = sum + value
		mul = mul * value
	}
	return mul - sum
}
