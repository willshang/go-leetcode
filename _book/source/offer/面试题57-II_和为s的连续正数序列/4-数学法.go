package main

import (
	"fmt"
)

func main() {
	fmt.Println(findContinuousSequence(9))
}

// target = nA1 + n(n-1)/2
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for i := (target + 1) / 2; i >= 2; i-- {
		nA1 := target - i*(i-1)/2
		if nA1 <= 0 {
			continue
		}
		if nA1%i == 0 {
			start := nA1 / i
			arr := make([]int, 0)
			for j := 0; j < i; j++ {
				arr = append(arr, start+j)
			}
			res = append(res, arr)
		}
	}
	return res
}
