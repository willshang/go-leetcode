package main

import (
	"fmt"
)

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{2, 3, 6, 9, 10}))
}

func canMakeArithmeticProgression(arr []int) bool {
	min, max := arr[0], arr[0]
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
		}
		if arr[i] >= max {
			max = arr[i]
		}
		m[arr[i]] = true
	}
	diff := (max - min) / (len(arr) - 1)
	for i := 0; i < len(m); i++ {
		value := min + diff*i
		if _, ok := m[value]; !ok {
			return false
		}
	}
	return true
}
