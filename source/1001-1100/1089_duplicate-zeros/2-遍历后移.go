package main

import "fmt"

func main() {
	//arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr := []int{0, 1, 7, 6, 0, 2, 0, 7}
	duplicateZeros(arr)
	fmt.Println(arr)
}

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}
}
