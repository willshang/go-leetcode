package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(largestPalindrome(2))
}

func makePalindrome(num int) int {
	numStr := strconv.Itoa(num)
	numBytes := []byte(numStr)
	palindrome := []byte(numStr)
	digitsCount := len(numBytes)
	for i := digitsCount - 1; i > -1; i-- {
		palindrome = append(palindrome, numBytes[i])
	}

	palindromNumStr := string(palindrome)
	palindromeNum, _ := strconv.Atoi(palindromNumStr)
	return palindromeNum
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	last := int(math.Pow10(n)) - 1
	first := last / 10
	for i := last; i > first; i-- {
		palindrome := makePalindrome(i)
		for j := last; j > first && palindrome < j*j; j-- {
			if palindrome%j == 0 {
				return palindrome % 1337
			}
		}
	}
	return 0
}

/*var res = []int{9, 987, 123, 597, 677, 1218, 877, 475}

func largestPalindrome(n int) int {
	return res[n-1]
}*/
