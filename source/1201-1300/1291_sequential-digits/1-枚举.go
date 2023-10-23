package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sequentialDigits(100, 300))
}

func sequentialDigits(low int, high int) []int {
	res := make([]int, 0)
	for i := 1; i <= 9; i++ {
		num := i
		for j := i + 1; j <= 9; j++ {
			num = num*10 + j
			if num >= low && num <= high {
				res = append(res, num)
			}
		}
	}
	sort.Ints(res)
	return res
}
