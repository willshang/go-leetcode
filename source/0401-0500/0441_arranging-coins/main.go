package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(arrangeCoins(9))
}

/* 数学推导
   (1+k)*k/2 = n
   k+k*k = 2*n
   k*k + k + 0.25 = 2*n + 0.25
   (k + 0.5) ^ 2 = 2*n +0.25
   k + 0.5 = sqrt(2*n + 0.25)
   k = sqrt(2*n + 0.25) - 0.5
*/
func arrangeCoins(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}
