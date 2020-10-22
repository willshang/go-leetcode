package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}

func search(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return true
		}
	}
	return false
}
