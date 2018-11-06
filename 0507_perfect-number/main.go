package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(24))
	fmt.Println(checkPerfectNumber(28))
}
func checkPerfectNumber(num int) bool {
	if num == 1{
		return false
	}

	sum := 1
	root := int(math.Sqrt(float64(num)))
	for i := 2; i <= root; i++{
		if num % i == 0{
			sum = sum + i + (num / i)
		}
	}
	return sum == num
}