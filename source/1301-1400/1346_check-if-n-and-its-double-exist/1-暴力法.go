package main

import "fmt"

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
}

// leetcode1346_检查整数及其两倍数是否存在
func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]*2 == arr[j] || arr[j]*2 == arr[i] {
				return true
			}
		}
	}
	return false
}
