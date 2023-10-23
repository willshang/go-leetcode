package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	list := make([]int, 0)
	list = append(list, n)
	level := 0
	for len(list) > 0 {
		level++
		length := len(list)
		for i := 0; i < length; i++ {
			value := list[i]
			for j := 1; j*j <= value; j++ {
				if j*j == value {
					return level
				}
				list = append(list, value-j*j)
			}
		}
		list = list[length:]
	}
	return level
}
