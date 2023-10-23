package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

// leetcode202_快乐数
func isHappy(n int) bool {
	now, next := n, nextValue(n)
	for now != next {
		now = nextValue(now)
		next = nextValue(nextValue(next))
	}
	if now == 1 {
		return true
	}
	return false
}

func nextValue(n int) int {
	ret := 0
	for n != 0 {
		ret = ret + (n%10)*(n%10)
		n = n / 10
	}
	return ret
}
