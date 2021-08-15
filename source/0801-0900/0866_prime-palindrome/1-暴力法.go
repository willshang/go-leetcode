package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(primePalindrome(1))
	fmt.Println(primePalindrome(2))
}

// leetcode866_回文素数
func primePalindrome(N int) int {
	if 8 <= N && N <= 11 { // 特判：偶数位的回文数都可以被11整除
		return 11
	}
	for i := 1; i < 100000; i++ {
		s := strconv.Itoa(i)
		temp := []byte(s)
		for i := 0; i < len(s)/2; i++ {
			temp[i], temp[len(s)-1-i] = temp[len(s)-1-i], temp[i]
		}
		target, _ := strconv.Atoi(s + string(temp[1:]))
		if target >= N && isPrime(target) {
			return target
		}
	}
	return -1
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
