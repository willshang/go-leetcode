package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

func exchange(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			res = append(res, nums[i])
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			res = append(res, nums[i])
		}
	}
	return res
}
