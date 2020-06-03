package main

import "fmt"

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
}

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left] < numbers[right] {
			return numbers[left]
		}
		mid := left + (right-left)/2
		if numbers[mid] > numbers[left] {
			left = mid + 1
		} else if numbers[mid] < numbers[left] {
			right = mid
		} else {
			left++
		}
	}
	return numbers[left]
}
