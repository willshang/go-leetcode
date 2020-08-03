package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		value := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				value = value * nums[j]
			}
		}
		res = append(res, value)
	}
	return res
}
