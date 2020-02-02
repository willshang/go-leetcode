package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k = k % n
	}
	if k == 0 || k == n {
		return
	}

	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[(i+k)%len(nums)] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = arr[i]
	}
}
