package main

import "fmt"

func main() {
	//arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr := []int{0, 1, 7, 6, 0, 2, 0, 7}
	duplicateZeros(arr)
	fmt.Println(arr)
}

// leetcode1089_复写零
func duplicateZeros(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count++
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			count--
			if i+count < len(arr) {
				arr[i+count] = 0
			}
			if i+count+1 < len(arr) {
				arr[i+count+1] = 0
			}
		} else if i+count < len(arr) {
			arr[i+count] = arr[i]
		}
	}
}
