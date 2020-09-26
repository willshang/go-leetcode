package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(10))
}

func fib(N int) int {
	temp1 := (1 + math.Sqrt(5)) / 2
	temp2 := (1 - math.Sqrt(5)) / 2
	fn := math.Round((math.Pow(temp1, float64(N)) - math.Pow(temp2, float64(N))) / math.Sqrt(5))
	return int(fn)
}
