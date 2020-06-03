package main

import "fmt"

func main() {
	//fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{1}))
	//fmt.Println(minArray([]int{3, 1, 3}))
	//fmt.Println(minArray([]int{10, 1, 10, 10, 10}))
}

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	mid := left
	for numbers[left] >= numbers[right] {
		if right-left == 1 {
			mid = right
			break
		}
		mid = (left + right) / 2
		if numbers[left] == numbers[right] && numbers[mid] == numbers[left] {
			return minInorder(numbers, left, right)
		}
		if numbers[mid] >= numbers[left] {
			left = mid
		} else if numbers[mid] <= numbers[right] {
			right = mid
		}
	}
	return numbers[mid]
}

func minInorder(numbers []int, left, right int) int {
	result := numbers[left]
	for i := left + 1; i <= right; i++ {
		if result > numbers[i] {
			result = numbers[i]
		}
	}
	return result
}
