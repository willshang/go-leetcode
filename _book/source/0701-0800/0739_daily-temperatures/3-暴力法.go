package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperatures([]int{73, 72, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(T []int) []int {
	j := 0
	for i := 0; i < len(T); i++ {
		for j = i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				T[i] = j - i
				break
			}
		}
		if j == len(T) {
			T[i] = 0
		}
	}
	return T
}
