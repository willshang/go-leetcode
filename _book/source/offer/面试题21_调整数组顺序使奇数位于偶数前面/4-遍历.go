package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

func exchange(nums []int) []int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
	}
	return nums
}
