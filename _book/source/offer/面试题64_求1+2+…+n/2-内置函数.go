package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumNums(10))
}

func sumNums(n int) int {
	return (int(math.Pow(float64(n), float64(2))) + n) >> 1
}
