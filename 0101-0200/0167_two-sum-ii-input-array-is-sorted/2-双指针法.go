package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

// leetcode 167 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	first := 0
	last := len(numbers) - 1
	result := make([]int, 2)

	for {
		if numbers[first]+numbers[last] == target {
			result[0] = first + 1
			result[1] = last + 1
			return result
		} else if numbers[first]+numbers[last] > target {
			last--
		} else {
			first++
		}
	}
}
