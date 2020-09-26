package main

import "fmt"

func main() {
	//arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	arr := []int{0, 1, 7, 6, 0, 2, 0, 7}
	duplicateZeros(arr)
	fmt.Println(arr)
}

func duplicateZeros(arr []int) {
	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			newArr = append(newArr, 0)
		}
		newArr = append(newArr, arr[i])
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = newArr[i]
	}
}
