package main

import "fmt"

func main() {
	fmt.Println(countDigitOne(12))
}

func countDigitOne(n int) int {
	res := 0
	digit := 1
	high := n / 10
	cur := n % 10
	low := 0
	for high != 0 || cur != 0 {
		if cur == 0 {
			res = res + high*digit
		} else if cur == 1 {
			res = res + high*digit + low + 1
		} else {
			res = res + (high+1)*digit
		}
		low = low + cur*digit
		cur = high % 10
		high = high / 10
		digit = digit * 10
	}
	return res
}
