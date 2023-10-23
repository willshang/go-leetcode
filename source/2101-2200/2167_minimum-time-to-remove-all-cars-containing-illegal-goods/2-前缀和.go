package main

import "fmt"

func main() {
	fmt.Println(minimumTime("1100101"))
	fmt.Println(minimumTime("101"))
}

func minimumTime(s string) int {
	n := len(s)
	res := n
	// left + middle(left,right)*2 + right
	// => i + (n-1-j) + 2*countOne(left,right)
	// => i + (n-1-j) + 2 *(pre[j+1] - pre[i])
	// => (i-2*pre[i]) + (2*pre[j+1]-j) + (n-1)
	// 求最小值
	a := 0
	sum := 0
	for j := 0; j < n; j++ {
		a = min(a, j-2*sum)
		sum = sum + int(s[j]-'0')
		b := 2*sum - j
		res = min(res, a+b+n-1)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
