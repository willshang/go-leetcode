package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
}

func findLucky(arr []int) int {
	res := make([]int, 501)
	for i := 0; i < len(arr); i++ {
		res[arr[i]]++
	}
	for i := 500; i >= 1; i-- {
		if res[i] == i {
			return i
		}
	}
	return -1
}
