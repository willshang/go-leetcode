package main

import "fmt"

func main() {
	for i := 0; i <= 22; i++ {
		fmt.Println(i, numberOf2sInRange(i))
	}
}

func numberOf2sInRange(n int) int {
	res := 0
	digit := 1
	high := n / 10
	cur := n % 10
	low := 0
	for high != 0 || cur != 0 {
		if cur > 2 {
			res = res + (high+1)*digit
		} else if cur == 2 {
			res = res + high*digit + low + 1
		} else {
			res = res + high*digit
		}
		low = low + cur*digit
		cur = high % 10
		high = high / 10
		digit = digit * 10
	}
	return res
}
