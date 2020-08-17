package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		max := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	return arr
}
