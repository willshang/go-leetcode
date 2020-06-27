package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func replaceElements(arr []int) []int {
	res := make([]int, len(arr))
	res[len(res)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i+1] > res[i+1] {
			res[i] = arr[i+1]
		} else {
			res[i] = res[i+1]
		}
	}
	return res
}
