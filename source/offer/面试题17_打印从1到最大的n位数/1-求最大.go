package main

import "fmt"

func main() {
	fmt.Println(printNumbers(3))
}

// 剑指offer_面试题17_打印从1到最大的n位数
func printNumbers(n int) []int {
	res := make([]int, 0)
	maxValue := 0
	for n > 0 {
		maxValue = maxValue*10 + 9
		n--
	}
	for i := 1; i <= maxValue; i++ {
		res = append(res, i)
	}
	return res
}
