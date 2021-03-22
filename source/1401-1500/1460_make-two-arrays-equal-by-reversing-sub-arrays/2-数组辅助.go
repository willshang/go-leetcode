package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}

func canBeEqual(target []int, arr []int) bool {
	temp := make([]int, 1001)
	for i := 0; i < len(target); i++ {
		temp[target[i]]++
		temp[arr[i]]--
	}
	for i := 0; i < len(temp); i++ {
		if temp[i] != 0 {
			return false
		}
	}
	return true
}
