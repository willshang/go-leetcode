package main

import "fmt"

func main() {
	//fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
}

func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if m[arr[i]*2] > 0 || (arr[i]%2 == 0 && m[arr[i]/2] > 0) {
			return true
		}
		m[arr[i]] = 1
	}
	return false
}
