package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCount([]int{4, 2, 1}))
}

func minCount(coins []int) int {
	res := 0
	for i := 0; i < len(coins); i++ {
		res = res + int(math.Ceil(float64(coins[i])/2))
	}
	return res
}
