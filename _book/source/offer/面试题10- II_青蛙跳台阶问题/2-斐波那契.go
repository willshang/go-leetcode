package main

import "fmt"

func main() {
	fmt.Println(numWays(100))
}

// 剑指offer_面试题10-II.青蛙跳台阶问题
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := (first + second) % 1000000007
		first = second
		second = third
	}
	return second
}
