package main

import "fmt"

func main() {
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(7))
	fmt.Println(isUgly(9))
	fmt.Println(isUgly(11))
}

// leetcode263_丑数
func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else if num%3 == 0 {
			num = num / 3
		} else if num%5 == 0 {
			num = num / 5
		} else {
			return false
		}
	}
	return true
}
