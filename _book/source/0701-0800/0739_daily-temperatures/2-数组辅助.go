package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(dailyTemperatures([]int{73, 72, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	arr := make([]int, 101)
	for i := 0; i < len(arr); i++ {
		arr[i] = math.MaxInt64
	}
	for i := len(T) - 1; i >= 0; i-- {
		temp := math.MaxInt64
		for t := T[i] + 1; t < 101; t++ {
			if arr[t] < temp {
				temp = arr[t]
			}
		}
		if temp < math.MaxInt64 {
			res[i] = temp - i
		}
		arr[T[i]] = i
	}
	return res
}
