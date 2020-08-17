package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 7, 10}))
}

func findSpecialInteger(arr []int) int {
	count := 1
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
			if count > len(arr)/4 {
				return arr[i]
			}
		} else {
			res = arr[i]
			count = 1
		}
	}
	return res
}
