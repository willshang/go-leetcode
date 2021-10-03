package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(7, -3))
}

// 剑指OfferII001.整数除法
func divide(a int, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	if b == 1 {
		return a
	}
	flag, count := 1, 1
	if a < 0 {
		flag = -flag
		a = -a
	}
	if b < 0 {
		flag = -flag
		b = -b
	}
	x, y, z := a, b, 0
	temp := y
	for x-y >= 0 {
		for x-y >= 0 {
			x = x - y
			z = z + count
			y = y + y
			count = count + count
		}
		y = temp
		count = 1
	}
	if z > math.MaxInt32 {
		return math.MaxInt32
	}
	if flag < 0 {
		return -z
	}
	return z
}
