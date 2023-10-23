package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

// 剑指OfferII006.排序数组中两个数字之和
func twoSum(numbers []int, target int) []int {
	first := 0
	last := len(numbers) - 1
	result := make([]int, 2)

	for {
		if numbers[first]+numbers[last] == target {
			result[0] = first
			result[1] = last
			return result
		} else if numbers[first]+numbers[last] > target {
			last--
		} else {
			first++
		}
	}
}
