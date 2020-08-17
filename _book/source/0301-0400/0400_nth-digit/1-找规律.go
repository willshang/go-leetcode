package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNthDigit(9999))
	fmt.Println(findNthDigit(13))
	fmt.Println(findNthDigit(11))
}

func findNthDigit(n int) int {
	if n < 0 {
		return -1
	}
	digits := 1
	for {
		numbers := countOfIntegers(digits)
		if n < numbers*digits {
			return digitAtIndex(n, digits)
		}
		n = n - numbers*digits
		digits++
	}
}

func countOfIntegers(n int) int {
	if n == 1 {
		return 10
	}
	count := math.Pow(float64(10), float64(n-1))
	return 9 * int(count)
}

func digitAtIndex(n, digits int) int {
	number := beginNumber(digits) + n/digits
	indexFromRight := digits - n%digits
	for i := 1; i < indexFromRight; i++ {
		number = number / 10
	}
	return number % 10
}

func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow(float64(10), float64(digits-1)))
}
