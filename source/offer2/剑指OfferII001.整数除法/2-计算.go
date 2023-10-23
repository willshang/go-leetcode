package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(7, -3))
}

func divide(a int, b int) int {
	res := a / b
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
