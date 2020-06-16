package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(5))
}

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		current := i*i + j*j
		if current < c {
			i++
		} else if current > c {
			j--
		} else {
			return true
		}
	}
	return false
}
