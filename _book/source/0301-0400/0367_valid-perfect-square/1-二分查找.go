package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(100))
	fmt.Println(isPerfectSquare(99))
}

// leetcode367_有效的完全平方数
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left := 2
	right := num / 2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
