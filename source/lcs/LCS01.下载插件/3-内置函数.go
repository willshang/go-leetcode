package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(leastMinutes(8))
}

func leastMinutes(n int) int {
	res := int(math.Ceil(math.Log(float64(n)) / math.Log(2)))
	return res + 1
}
