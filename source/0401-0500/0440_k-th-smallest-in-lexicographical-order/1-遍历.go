package main

import "fmt"

func main() {
	fmt.Println(findKthNumber(10, 3))
}

// leetcode440_字典序的第K小数字
func findKthNumber(n int, k int) int {
	pre := 1
	count := 1
	for count < k {
		total := getCount(pre, n)
		if count+total > k { // 超过了，*10继续尝试，缩小范围
			pre = pre * 10
			count++
		} else if count+total <= k { // 小了，尝试下一个数字开头，扩大范围
			pre++
			count = count + total
		}
	}
	return pre
}

// 以数字i开头且不超过最大数字n的数字个数
func getCount(pre, n int) int {
	count := 0
	a := pre
	b := pre + 1
	for a <= n {
		if b < n+1 {
			count = count + b - a
		} else {
			count = count + n + 1 - a
		}
		a = a * 10
		b = b * 10
	}
	return count
}
