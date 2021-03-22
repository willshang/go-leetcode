package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}

func findKthPositive(arr []int, k int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i]-(i+1) >= k {
			return k + i
		}
	}
	return k + len(arr)
}
