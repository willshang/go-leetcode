package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(7, -3))
}

// leetcode29_两数相除
func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	flag, count := 1, 1
	if dividend < 0 {
		flag = -flag
		dividend = -dividend
	}
	if divisor < 0 {
		flag = -flag
		divisor = -divisor
	}
	a, b, c := dividend, divisor, 0
	temp := b
	for a-b >= 0 {
		for a-b >= 0 {
			a = a - b
			c = c + count
			b = b + b
			count = count + count
		}
		b = temp
		count = 1
	}
	if c > math.MaxInt32 {
		return math.MaxInt32
	}
	if flag < 0 {
		return -c
	}
	return c
}
