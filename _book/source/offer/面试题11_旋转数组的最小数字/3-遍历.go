package main

import (
	"fmt"
)

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
}

// 剑指offer_面试题11_旋转数组的最小数字
func minArray(numbers []int) int {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}
