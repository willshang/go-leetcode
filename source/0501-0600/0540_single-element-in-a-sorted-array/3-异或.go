package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 3, 3, 4, 4, 8, 8, 9}))
}

func singleNonDuplicate(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
