package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(12321))
}

//执行用时：68 ms
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x = x / 10
	}
	//for example:
	// x = 1221  => x = 12 revertedNumber = 12
	// x = 12321 => x = 12 revertedNumber = 123
	return x == revertedNumber || x == revertedNumber/10
}

//执行用时：68 ms
/*func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}*/
