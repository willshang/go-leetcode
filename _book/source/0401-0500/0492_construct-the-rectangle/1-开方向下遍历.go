package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(5))
}

// leetcode492_构造矩形
func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i > 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{area, 1}
}
