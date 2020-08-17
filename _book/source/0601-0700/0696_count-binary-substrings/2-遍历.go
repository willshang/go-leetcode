package main

import (
	"fmt"
)

func main() {
	str := "00110011"
	fmt.Println(countBinarySubstrings(str))
}

func countBinarySubstrings(s string) int {
	res := 0
	arr := make([]int, 0)
	arr = append(arr, 1)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			arr[len(arr)-1]++
		} else {
			arr = append(arr, 1)
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			res = res + arr[i+1]
		} else {
			res = res + arr[i]
		}
	}
	return res
}
