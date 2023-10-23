package main

import "fmt"

func main() {
	fmt.Println(threeConsecutiveOdds([]int{1, 3, 5, 7, 2}))
}

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
			return true
		}
	}
	return false
}
