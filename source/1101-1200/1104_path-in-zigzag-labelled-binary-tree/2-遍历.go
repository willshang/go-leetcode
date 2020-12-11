package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(pathInZigZagTree(7))
	//fmt.Println(pathInZigZagTree(8))
	//fmt.Println(pathInZigZagTree(9))
	//fmt.Println(pathInZigZagTree(15))
	fmt.Println(pathInZigZagTree(14))
}

func pathInZigZagTree(label int) []int {
	length := int(math.Log2(float64(label)))
	res := make([]int, length+1)
	res[length] = label
	length--
	i := 1
	for length >= 0 {
		target := int(math.Pow(2, float64(length+1))) + int(math.Pow(2, float64(length))) - 1
		if i%2 == 1 {
			res[length] = target - label/2
		} else {
			res[length] = label / 2
		}
		i++
		length--
		label = label / 2
	}
	return res
}
