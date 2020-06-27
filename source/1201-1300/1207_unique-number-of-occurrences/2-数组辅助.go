package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}

func uniqueOccurrences(arr []int) bool {
	tempArr := make([]int, 2001)
	for _, v := range arr {
		tempArr[v+1000]++
	}
	temp := make(map[int]bool)
	for _, v := range tempArr {
		if v == 0 {
			continue
		}
		if temp[v] == true {
			return false
		}
		temp[v] = true
	}
	return true
}
