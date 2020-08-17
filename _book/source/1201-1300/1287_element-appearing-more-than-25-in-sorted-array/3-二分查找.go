package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 7, 10}))
	//fmt.Println(findSpecialInteger([]int{1}))
}

func findSpecialInteger(arr []int) int {
	length := len(arr) / 4
	for i := 1; i <= 2; i++ {
		value := arr[length*i]
		left := leftSearch(arr, value)
		right := rightSearch(arr, value)
		if right-left+1 > length {
			return value
		}
	}
	return arr[3*length]
}

func leftSearch(arr []int, value int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= value {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func rightSearch(arr []int, value int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > value {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right - 1
}
