package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNoZeroIntegers(11))
}

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		if contains(i) || contains(n-i) {
			continue
		}
		return []int{i, n - i}
	}
	return nil
}

func contains(num int) bool {
	for num > 0 {
		if num%10 == 0 {
			return false
		}
		num = num / 10
	}
	return true
}
